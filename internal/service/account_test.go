package service

import (
	"fmt"
	"testing"

	"github.com/1612180/chat_stranger/internal/mock/mock_repository"
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/config"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewAccountService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mAccountRepo := mock_repository.NewMockAccountRepo(ctrl)
	cfg := config.NewConfig(variable.Test)

	assert.Equal(t, &defaultAccountService{accountRepo: mAccountRepo, cfg: cfg}, NewAccountService(mAccountRepo, cfg))
}

func TestDefaultAccountService_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mAccountRepo := mock_repository.NewMockAccountRepo(ctrl)
	cfg := config.NewConfig(variable.Test)
	accountService := defaultAccountService{accountRepo: mAccountRepo, cfg: cfg}

	mAccountRepo.EXPECT().
		CreateUserCredential("a", "b", gomock.Any()).
		Return(model.User{}, model.Credential{}, nil)
	mAccountRepo.EXPECT().
		CreateUserCredential("c", "d", gomock.Any()).
		Return(model.User{}, model.Credential{}, fmt.Errorf("err"))

	testCases := []struct {
		inShowName string
		inRegName  string
		inPassword string
		ok         bool
	}{
		{
			inShowName: "a",
			inRegName:  "b",
			ok:         true,
		},
		{
			inShowName: "c",
			inRegName:  "d",
			ok:         false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			_, err := accountService.SignUp(tc.inShowName, tc.inRegName, tc.inPassword)
			if tc.ok {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDefaultAccountService_LogIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mAccountRepo := mock_repository.NewMockAccountRepo(ctrl)
	cfg := config.NewConfig(variable.Test)
	accountService := defaultAccountService{accountRepo: mAccountRepo, cfg: cfg}

	hashedPass1, err := bcrypt.GenerateFromPassword([]byte("pass"), bcrypt.DefaultCost)
	if err != nil {
		t.Error(err)
	}

	mAccountRepo.EXPECT().
		FindUserCredentialByRegName("name").
		Return(model.User{}, model.Credential{HashedPassword: string(hashedPass1)}, nil)
	mAccountRepo.EXPECT().
		FindUserCredentialByRegName("name2").
		Return(model.User{}, model.Credential{}, nil)
	mAccountRepo.EXPECT().
		FindUserCredentialByRegName("name3").
		Return(model.User{}, model.Credential{}, fmt.Errorf("err"))

	testCases := []struct {
		inRegName  string
		inPassword string
		ok         bool
	}{
		{
			inRegName:  "name",
			inPassword: "pass",
			ok:         true,
		},
		{
			inRegName: "name2",
			ok:        false,
		},
		{
			inRegName: "name3",
			ok:        false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			_, err := accountService.LogIn(tc.inRegName, tc.inPassword)
			if tc.ok {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDefaultAccountService_Info(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mAccountRepo := mock_repository.NewMockAccountRepo(ctrl)
	cfg := config.NewConfig(variable.Test)
	accountService := defaultAccountService{accountRepo: mAccountRepo, cfg: cfg}

	mAccountRepo.EXPECT().
		FindUserCredential(1).
		Return(model.User{ID: 1, ShowName: "a"}, model.Credential{}, nil)
	mAccountRepo.EXPECT().
		FindUserCredential(2).
		Return(model.User{}, model.Credential{}, fmt.Errorf("err"))

	testCases := []struct {
		inUserID int
		outUser  model.User
		ok       bool
	}{
		{
			inUserID: 1,
			outUser:  model.User{ID: 1, ShowName: "a"},
			ok:       true,
		},
		{
			inUserID: 2,
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			user, err := accountService.Info(tc.inUserID)
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, tc.outUser, user)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDefaultAccountService_UpdateInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mAccountRepo := mock_repository.NewMockAccountRepo(ctrl)
	cfg := config.NewConfig(variable.Test)
	accountService := defaultAccountService{accountRepo: mAccountRepo, cfg: cfg}

	mAccountRepo.EXPECT().
		UpdateUser(1, "name", "gender", 2000).
		Return(model.User{ID: 1, ShowName: "name", Gender: "gender", BirthYear: 2000}, nil)
	mAccountRepo.EXPECT().
		UpdateUser(2, "name2", "gender2", 2001).
		Return(model.User{}, fmt.Errorf("err"))

	testCases := []struct {
		inUserID    int
		inShowName  string
		inGender    string
		inBirthYear int
		outUser     model.User
		ok          bool
	}{
		{
			inUserID:    1,
			inShowName:  "name",
			inGender:    "gender",
			inBirthYear: 2000,
			outUser:     model.User{ID: 1, ShowName: "name", Gender: "gender", BirthYear: 2000},
			ok:          true,
		},
		{
			inUserID:    2,
			inShowName:  "name2",
			inGender:    "gender2",
			inBirthYear: 2001,
			ok:          false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			user, err := accountService.UpdateInfo(tc.inUserID, tc.inShowName, tc.inGender, tc.inBirthYear)
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, tc.outUser, user)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
