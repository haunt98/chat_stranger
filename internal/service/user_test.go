package service

import (
	"testing"

	"github.com/1612180/chat_stranger/internal/mock/mock_repository"
	"github.com/1612180/chat_stranger/internal/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_repository.NewMockUserRepository(ctrl)
	m.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(true)

	userService := userService{
		userRepo: m,
	}

	ok := userService.SignUp(&model.User{})
	assert.True(t, ok)
}

// https://play.golang.org/p/p8GiKu1Ys75

func TestUserService_LogIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("false", func(t *testing.T) {
		m := mock_repository.NewMockUserRepository(ctrl)
		m.EXPECT().
			FindByRegisterName(gomock.Any()).
			Return(
				&model.User{},
				&model.Credential{},
				true,
			)

		userService := userService{
			userRepo: m,
		}

		ok := userService.LogIn(&model.User{})
		assert.False(t, ok)
	})

	t.Run("true", func(t *testing.T) {
		m := mock_repository.NewMockUserRepository(ctrl)
		m.EXPECT().
			FindByRegisterName(gomock.Any()).
			Return(
				&model.User{},
				&model.Credential{HashedPassword: "$2a$10$d8Eaak/7DcJp06A2dBhql.NNWFnFNKBWyCOyiv/bVk/wl6tpwD/pO"},
				true,
			)

		userService := userService{
			userRepo: m,
		}

		ok := userService.LogIn(&model.User{Password: "a"})
		assert.True(t, ok)
	})
}

func TestUserService_UpdateInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("false", func(t *testing.T) {
		m := mock_repository.NewMockUserRepository(ctrl)
		m.EXPECT().
			Find(gomock.Any()).
			Return(&model.User{}, &model.Credential{}, false)

		userService := userService{
			userRepo: m,
		}

		ok := userService.UpdateInfo(1, &model.User{})
		assert.False(t, ok)
	})

	t.Run("false", func(t *testing.T) {
		m := mock_repository.NewMockUserRepository(ctrl)
		m.EXPECT().
			Find(gomock.Any()).
			Return(&model.User{}, &model.Credential{}, true)
		m.EXPECT().
			UpdateInfo(gomock.Any(), gomock.Any()).
			Return(false)

		userService := userService{
			userRepo: m,
		}

		ok := userService.UpdateInfo(1, &model.User{})
		assert.False(t, ok)
	})

	t.Run("true", func(t *testing.T) {
		m := mock_repository.NewMockUserRepository(ctrl)
		m.EXPECT().
			Find(gomock.Any()).
			Return(&model.User{}, &model.Credential{}, true)
		m.EXPECT().
			UpdateInfo(gomock.Any(), gomock.Any()).
			Return(true)

		userService := userService{
			userRepo: m,
		}

		ok := userService.UpdateInfo(1, &model.User{})
		assert.True(t, ok)
	})
}

func TestUserService_UpdatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("false", func(t *testing.T) {
		m := mock_repository.NewMockUserRepository(ctrl)
		m.EXPECT().
			UpdatePassword(gomock.Any(), gomock.Any()).
			Return(false)

		userService := userService{
			userRepo: m,
		}

		ok := userService.UpdatePassword(1, &model.User{})
		assert.False(t, ok)
	})

	t.Run("true", func(t *testing.T) {
		m := mock_repository.NewMockUserRepository(ctrl)
		m.EXPECT().
			UpdatePassword(gomock.Any(), gomock.Any()).
			Return(true)

		userService := userService{
			userRepo: m,
		}

		ok := userService.UpdatePassword(1, &model.User{})
		assert.True(t, ok)
	})
}
