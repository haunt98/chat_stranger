package handler

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/1612180/chat_stranger/internal/mock/mock_service"
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/config"
	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/1612180/chat_stranger/internal/pkg/token"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/dghubble/sling"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewAccountHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mAccountService := mock_service.NewMockAccountService(ctrl)

	assert.Equal(t, &AccountHandler{accountService: mAccountService}, NewAccountHandler(mAccountService))
}

func TestAccountHandler_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mAccountService := mock_service.NewMockAccountService(ctrl)
	cfg := config.NewConfig(variable.Test)

	accountHandler := AccountHandler{accountService: mAccountService}
	router := NewRouter(&accountHandler, nil, cfg)

	mAccountService.EXPECT().
		SignUp("a", "reg", "pass").
		Return("tkn", nil)
	mAccountService.EXPECT().
		SignUp("b", "reg2", "pass2").
		Return("", fmt.Errorf("err"))

	testCases := []struct {
		inSubmit    interface{}
		outHTTPCode int
		outResponse response.Response
	}{
		{
			inSubmit:    SignUpSubmit{ShowName: "a", RegisterName: "reg", Password: "pass"},
			outHTTPCode: 200,
			outResponse: response.CreateWithData(100, "tkn"),
		},
		{
			inSubmit:    SignUpSubmit{ShowName: "b", RegisterName: "reg2", Password: "pass2"},
			outHTTPCode: 500,
			outResponse: response.Create(101),
		},
		{
			inSubmit:    SignUpSubmit{ShowName: "", RegisterName: "", Password: ""},
			outHTTPCode: 400,
			outResponse: response.Create(103),
		},
		{
			inSubmit:    nil,
			outHTTPCode: 400,
			outResponse: response.Create(102),
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			req, err := sling.New().Post(variable.APIPrefix + "/auth/signup").BodyJSON(tc.inSubmit).Request()
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.outHTTPCode, w.Code)

			var res response.Response
			if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
				t.Error(err)
			}
			assert.Equal(t, tc.outResponse.Code, res.Code)
			if res.Code == 100 {
				tkn, ok := tc.outResponse.Data.(string)
				assert.True(t, ok)
				assert.Equal(t, tkn, res.Data)
			}
		})
	}
}

func TestAccountHandler_LogIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mAccountService := mock_service.NewMockAccountService(ctrl)
	cfg := config.NewConfig(variable.Test)

	accountHandler := AccountHandler{accountService: mAccountService}
	router := NewRouter(&accountHandler, nil, cfg)

	mAccountService.EXPECT().
		LogIn("reg", "pass").
		Return("tkn", nil)
	mAccountService.EXPECT().
		LogIn("reg2", "pass2").
		Return("", fmt.Errorf("err"))

	testCases := []struct {
		inSubmit    interface{}
		outHTTPCode int
		outResponse response.Response
	}{
		{
			inSubmit:    LogInSubmit{"reg", "pass"},
			outHTTPCode: 200,
			outResponse: response.CreateWithData(200, "tkn"),
		},
		{
			inSubmit:    LogInSubmit{"reg2", "pass2"},
			outHTTPCode: 500,
			outResponse: response.Create(201),
		},
		{
			inSubmit:    LogInSubmit{"", ""},
			outHTTPCode: 400,
			outResponse: response.Create(203),
		},
		{
			inSubmit:    nil,
			outHTTPCode: 400,
			outResponse: response.Create(202),
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("inSubmit=%+v", tc.inSubmit), func(t *testing.T) {
			req, err := sling.New().Post(variable.APIPrefix + "/auth/login").BodyJSON(tc.inSubmit).Request()
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.outHTTPCode, w.Code)

			var res response.Response
			if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
				t.Error(err)
			}
			assert.Equal(t, tc.outResponse.Code, res.Code)
			if res.Code == 200 {
				tkn, ok := tc.outResponse.Data.(string)
				assert.True(t, ok)
				assert.Equal(t, tkn, res.Data)
			}
		})
	}
}

func TestAccountHandler_Info(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mAccountService := mock_service.NewMockAccountService(ctrl)
	cfg := config.NewConfig(variable.Test)

	accountHandler := AccountHandler{accountService: mAccountService}
	router := NewRouter(&accountHandler, nil, cfg)

	mAccountService.EXPECT().
		Info(1).
		Return(model.User{ID: 1}, nil)
	mAccountService.EXPECT().
		Info(2).
		Return(model.User{}, fmt.Errorf("err"))

	tkn1, err := token.Create(token.AccountClaims{ID: 1, Role: "user"}, cfg.Get(variable.JWTSecret))
	if err != nil {
		t.Error(err)
	}
	tkn2, err := token.Create(token.AccountClaims{ID: 2, Role: "user"}, cfg.Get(variable.JWTSecret))
	if err != nil {
		t.Error(err)
	}

	testCases := []struct {
		inHeader    string
		outHTTPCode int
		outResponse response.Response
	}{
		{
			inHeader:    "Bearer " + tkn1,
			outHTTPCode: 200,
			outResponse: response.CreateWithData(300, model.User{ID: 1}),
		},
		{
			inHeader:    "Bearer " + tkn2,
			outHTTPCode: 500,
			outResponse: response.Create(301),
		},
		{
			inHeader:    "",
			outHTTPCode: 403,
			outResponse: response.Create(999),
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			req, err := sling.New().Set("Authorization", tc.inHeader).Get(variable.APIPrefix + "/me").Request()
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.outHTTPCode, w.Code)

			var res response.Response
			if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
				t.Error(err)
			}
			assert.Equal(t, tc.outResponse.Code, res.Code)
			if res.Code == 300 {
				user, ok := tc.outResponse.Data.(model.User)
				assert.True(t, ok)
				assert.Equal(t, tc.outResponse.Data, user)
			}
		})
	}
}

func TestAccountHandler_UpdateInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mAccountService := mock_service.NewMockAccountService(ctrl)
	cfg := config.NewConfig(variable.Test)

	accountHandler := AccountHandler{accountService: mAccountService}
	router := NewRouter(&accountHandler, nil, cfg)

	mAccountService.EXPECT().
		UpdateInfo(1, "name", "gender", 2000).
		Return(model.User{ID: 1, ShowName: "name", Gender: "gender", BirthYear: 2000}, nil)
	mAccountService.EXPECT().
		UpdateInfo(2, "name2", "gender2", 2001).
		Return(model.User{}, fmt.Errorf("err"))

	tkn1, err := token.Create(token.AccountClaims{ID: 1, Role: "user"}, cfg.Get(variable.JWTSecret))
	if err != nil {
		t.Error(err)
	}
	tkn2, err := token.Create(token.AccountClaims{ID: 2, Role: "user"}, cfg.Get(variable.JWTSecret))
	if err != nil {
		t.Error(err)
	}

	testCases := []struct {
		inHeader    string
		inSubmit    interface{}
		outHTTPCode int
		outResponse response.Response
	}{
		{
			inHeader:    "Bearer " + tkn1,
			inSubmit:    InfoSubmit{ShowName: "name", Gender: "gender", BirthYear: 2000},
			outHTTPCode: 200,
			outResponse: response.CreateWithData(120, model.User{ID: 1, ShowName: "name", Gender: "gender", BirthYear: 2000}),
		},
		{
			inHeader:    "Bearer " + tkn2,
			inSubmit:    InfoSubmit{ShowName: "name2", Gender: "gender2", BirthYear: 2001},
			outHTTPCode: 500,
			outResponse: response.Create(121),
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			req, err := sling.New().Set("Authorization", tc.inHeader).Put(variable.APIPrefix + "/me").BodyJSON(tc.inSubmit).Request()
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.outHTTPCode, w.Code)

			var res response.Response
			if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
				t.Error(err)
			}
			assert.Equal(t, tc.outResponse.Code, res.Code)
			if res.Code == 120 {
				user, ok := tc.outResponse.Data.(model.User)
				assert.True(t, ok)
				assert.Equal(t, tc.outResponse.Data, user)
			}
		})
	}
}
