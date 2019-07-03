package service

import (
	"testing"

	"github.com/1612180/chat_stranger/internal/mock/mock_repository"
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_repository.NewMockUserRepository(ctrl)
	m.
		EXPECT().
		Find(gomock.Eq(1)).
		Return(&model.User{ID: 1}, true)

	userService := UserService{
		userRepo: m,
	}

	user, ok := userService.Find(1)
	assert.Equal(t, true, ok)
	assert.Equal(t, &model.User{ID: 1}, user)
}
