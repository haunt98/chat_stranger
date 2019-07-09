package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1612180/chat_stranger/internal/model"

	"github.com/1612180/chat_stranger/internal/pkg/variable"

	"github.com/1612180/chat_stranger/internal/mock/mock_repository"
	"github.com/1612180/chat_stranger/internal/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserRepository(ctrl)
	mockRoomRepo := mock_repository.NewMockRoomRepository(ctrl)
	mockMemberRepo := mock_repository.NewMockMemberRepo(ctrl)
	mockMessageRepo := mock_repository.NewMockMessageRepo(ctrl)

	mockUserRepo.
		EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(true)

	userService := service.NewUserService(mockUserRepo)
	chatService := service.NewChatService(mockRoomRepo, mockMemberRepo, mockMessageRepo)
	userHandler := NewUserHandler(userService)
	chatHandler := NewChatHandler(chatService)
	router := NewRouter(userHandler, chatHandler, true)

	w := httptest.NewRecorder()

	// mock user repo => id is not changed when create
	jsonUser, err := json.Marshal(model.User{ID: 1})
	if err != nil {
		t.Error(err)
	}

	// send test request
	req, err := http.NewRequest("POST", variable.APIPrefix+"/auth/signup", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Error(err)
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var wbody interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &wbody); err != nil {
		t.Error(err)
	}
	res := wbody.(map[string]interface{})
	assert.Equal(t, 100, int(res["code"].(float64)))
}
