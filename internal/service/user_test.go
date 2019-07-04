package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/1612180/chat_stranger/internal/mock/mock_repository"
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/golang/mock/gomock"
)

func TestUserService_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_repository.NewMockUserRepository(ctrl)
	m.
		EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(true)

	userService := UserService{
		userRepo: m,
	}

	ok := userService.SignUp(&model.User{})
	assert.Equal(t, true, ok)
}

// https://play.golang.org/p/vmQwsmhKc24

func TestUserService_LogIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_repository.NewMockUserRepository(ctrl)
	m.
		EXPECT().
		FindByRegisterName(gomock.Any()).
		Return(
			&model.User{},
			&model.Credential{HashedPassword: "$2a$10$d8Eaak/7DcJp06A2dBhql.NNWFnFNKBWyCOyiv/bVk/wl6tpwD/pO"},
			true,
		)

	userService := UserService{
		userRepo: m,
	}

	ok := userService.LogIn(&model.User{Password: "a"})
	assert.Equal(t, true, ok)
}
