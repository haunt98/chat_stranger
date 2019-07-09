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

	m := mock_repository.NewMockUserRepository(ctrl)
	m.
		EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(true)

	userService := service.NewUserService(m)
	userHandler := NewUserHandler(userService)
	router := NewRouter(userHandler)

	w := httptest.NewRecorder()
	jsonUser, err := json.Marshal(model.User{})
	if err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest("POST", variable.APIPrefix+"/auth/signup", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Error(err)
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// TODO assert body
}
