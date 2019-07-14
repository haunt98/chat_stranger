package service

import (
	"testing"

	"github.com/1612180/chat_stranger/internal/mock/mock_repository"
	"github.com/1612180/chat_stranger/internal/model"
	"golang.org/x/crypto/bcrypt"

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
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("b"), bcrypt.DefaultCost)
		if err != nil {
			t.Error(err)
		}

		m := mock_repository.NewMockUserRepository(ctrl)
		m.EXPECT().
			FindByRegisterName("a").
			Return(
				&model.User{},
				&model.Credential{HashedPassword: string(hashedPassword)},
				true,
			)

		userService := userService{
			userRepo: m,
		}

		ok := userService.LogIn(&model.User{RegisterName: "a", Password: "b"})
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
