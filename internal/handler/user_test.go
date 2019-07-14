package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1612180/chat_stranger/internal/mock/mock_service"
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/jwt"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/spf13/viper"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("100", func(t *testing.T) {
		userService := mock_service.NewMockUserService(ctrl)
		chatService := mock_service.NewMockChatService(ctrl)

		userService.EXPECT().
			SignUp(gomock.Any()).
			Return(true)

		userHandler := NewUserHandler(userService)
		chatHandler := NewChatHandler(chatService)
		router := NewRouter(userHandler, chatHandler, true)

		w := httptest.NewRecorder()

		// mock request
		jsonUser, err := json.Marshal(model.User{})
		if err != nil {
			t.Error(err)
		}

		// send test request
		req, err := http.NewRequest("POST", variable.APIPrefix+"/auth/signup", bytes.NewBuffer(jsonUser))
		if err != nil {
			t.Error(err)
		}
		router.ServeHTTP(w, req)

		// check http code
		assert.Equal(t, 200, w.Code)

		// check response body
		var wbody interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &wbody); err != nil {
			t.Error(err)
		}
		res := wbody.(map[string]interface{})
		assert.EqualValues(t, 100, res["code"])
	})

	t.Run("101", func(t *testing.T) {
		userService := mock_service.NewMockUserService(ctrl)
		chatService := mock_service.NewMockChatService(ctrl)

		userService.EXPECT().
			SignUp(gomock.Any()).
			Return(false)

		userHandler := NewUserHandler(userService)
		chatHandler := NewChatHandler(chatService)
		router := NewRouter(userHandler, chatHandler, true)

		w := httptest.NewRecorder()

		// mock request
		jsonUser, err := json.Marshal(model.User{})
		if err != nil {
			t.Error(err)
		}

		// send test request
		req, err := http.NewRequest("POST", variable.APIPrefix+"/auth/signup", bytes.NewBuffer(jsonUser))
		if err != nil {
			t.Error(err)
		}
		router.ServeHTTP(w, req)

		// check http code
		assert.Equal(t, 200, w.Code)

		// check response body
		var wbody interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &wbody); err != nil {
			t.Error(err)
		}
		res := wbody.(map[string]interface{})
		assert.EqualValues(t, 101, res["code"])
	})

	t.Run("102", func(t *testing.T) {
		userService := mock_service.NewMockUserService(ctrl)
		chatService := mock_service.NewMockChatService(ctrl)

		userHandler := NewUserHandler(userService)
		chatHandler := NewChatHandler(chatService)
		router := NewRouter(userHandler, chatHandler, true)

		w := httptest.NewRecorder()

		// send test request
		req, err := http.NewRequest("POST", variable.APIPrefix+"/auth/signup", nil)
		if err != nil {
			t.Error(err)
		}
		router.ServeHTTP(w, req)

		// check http code
		assert.Equal(t, 200, w.Code)

		// check response body
		var wbody interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &wbody); err != nil {
			t.Error(err)
		}
		res := wbody.(map[string]interface{})
		assert.EqualValues(t, 102, res["code"])
	})
}

func TestUserHandler_LogIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("200", func(t *testing.T) {
		userService := mock_service.NewMockUserService(ctrl)
		chatService := mock_service.NewMockChatService(ctrl)

		userService.EXPECT().
			LogIn(gomock.Any()).
			Return(true)

		userHandler := NewUserHandler(userService)
		chatHandler := NewChatHandler(chatService)
		router := NewRouter(userHandler, chatHandler, true)

		w := httptest.NewRecorder()

		// mock request
		jsonUser, err := json.Marshal(model.User{})
		if err != nil {
			t.Error(err)
		}

		// send test request
		req, err := http.NewRequest("POST", variable.APIPrefix+"/auth/login", bytes.NewBuffer(jsonUser))
		if err != nil {
			t.Error(err)
		}
		router.ServeHTTP(w, req)

		// check http code
		assert.Equal(t, 200, w.Code)

		// check response body
		var wbody interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &wbody); err != nil {
			t.Error(err)
		}
		res := wbody.(map[string]interface{})
		assert.EqualValues(t, 200, res["code"])
	})

	t.Run("201", func(t *testing.T) {
		userService := mock_service.NewMockUserService(ctrl)
		chatService := mock_service.NewMockChatService(ctrl)

		userService.EXPECT().
			LogIn(gomock.Any()).
			Return(false)

		userHandler := NewUserHandler(userService)
		chatHandler := NewChatHandler(chatService)
		router := NewRouter(userHandler, chatHandler, true)

		w := httptest.NewRecorder()

		// mock request
		jsonUser, err := json.Marshal(model.User{})
		if err != nil {
			t.Error(err)
		}

		// send test request
		req, err := http.NewRequest("POST", variable.APIPrefix+"/auth/login", bytes.NewBuffer(jsonUser))
		if err != nil {
			t.Error(err)
		}
		router.ServeHTTP(w, req)

		// check http code
		assert.Equal(t, 200, w.Code)

		// check response body
		var wbody interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &wbody); err != nil {
			t.Error(err)
		}
		res := wbody.(map[string]interface{})
		assert.EqualValues(t, 201, res["code"])
	})

	t.Run("202", func(t *testing.T) {
		userService := mock_service.NewMockUserService(ctrl)
		chatService := mock_service.NewMockChatService(ctrl)

		userHandler := NewUserHandler(userService)
		chatHandler := NewChatHandler(chatService)
		router := NewRouter(userHandler, chatHandler, true)

		w := httptest.NewRecorder()

		// send test request
		req, err := http.NewRequest("POST", variable.APIPrefix+"/auth/login", nil)
		if err != nil {
			t.Error(err)
		}
		router.ServeHTTP(w, req)

		// check http code
		assert.Equal(t, 200, w.Code)

		// check response body
		var wbody interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &wbody); err != nil {
			t.Error(err)
		}
		res := wbody.(map[string]interface{})
		assert.EqualValues(t, 202, res["code"])
	})
}

func TestUserHandler_UpdateInfo(t *testing.T) {
	// Load config
	configutils.LoadConfiguration("chat_stranger", "config", "configs")
	t.Log("Bla bla", viper.GetString(variable.JWTSecret))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("999", func(t *testing.T) {
		userService := mock_service.NewMockUserService(ctrl)
		chatService := mock_service.NewMockChatService(ctrl)

		userHandler := NewUserHandler(userService)
		chatHandler := NewChatHandler(chatService)
		router := NewRouter(userHandler, chatHandler, true)

		w := httptest.NewRecorder()

		// send test request
		req, err := http.NewRequest("PUT", variable.APIPrefix+"/me", nil)
		if err != nil {
			t.Error(err)
		}
		router.ServeHTTP(w, req)

		// check http code
		assert.Equal(t, 403, w.Code)

		// check response body
		var wbody interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &wbody); err != nil {
			t.Error(err)
		}
		res := wbody.(map[string]interface{})
		assert.EqualValues(t, 999, res["code"])
	})

	t.Run("120", func(t *testing.T) {
		userService := mock_service.NewMockUserService(ctrl)
		chatService := mock_service.NewMockChatService(ctrl)

		userService.EXPECT().
			UpdateInfo(gomock.Any(), gomock.Any()).
			Return(true)

		userHandler := NewUserHandler(userService)
		chatHandler := NewChatHandler(chatService)
		router := NewRouter(userHandler, chatHandler, true)

		w := httptest.NewRecorder()

		// create token
		s, ok := jwt.Create(jwt.SignClaims{}, viper.GetString(variable.JWTSecret))
		t.Log("Bla bla", viper.GetString(variable.JWTSecret))

		assert.True(t, ok)

		// mock request
		jsonUser, err := json.Marshal(model.User{})
		if err != nil {
			t.Error(err)
		}

		// send test request
		req, err := http.NewRequest("PUT", variable.APIPrefix+"/me", bytes.NewBuffer(jsonUser))
		if err != nil {
			t.Error(err)
		}
		req.Header.Set("Authorization", "Bearer "+s)
		req.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		// check http code
		assert.Equal(t, 200, w.Code)

		// check response body
		var wbody interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &wbody); err != nil {
			t.Error(err)
		}
		res := wbody.(map[string]interface{})
		assert.EqualValues(t, 999, res["code"])
	})
}
