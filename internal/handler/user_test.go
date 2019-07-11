package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1612180/chat_stranger/internal/mock/mock_service"
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/variable"

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
}
