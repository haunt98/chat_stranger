package service

import (
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/repository"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

//go:generate $GOPATH/bin/mockgen -destination=../mock/mock_service/mock_user.go -source=user.go

type UserService interface {
	SignUp(user *model.User) bool
	LogIn(user *model.User) bool
	Info(id int) (*model.User, bool)
	UpdateInfo(id int, new *model.User) bool
	UpdatePassword(id int, new *model.User) bool
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// implement

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) SignUp(user *model.User) bool {
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "service",
			"target": "user",
			"action": "sign up",
		}).Error(err)
		return false
	}

	account := model.Credential{
		RegisterName:   user.RegisterName,
		HashedPassword: string(hashedPassword),
	}

	ok := s.userRepo.Create(user, &account)
	if !ok {
		return false
	}
	return true
}

func (s *userService) LogIn(user *model.User) bool {
	userInDB, credential, ok := s.userRepo.FindByRegisterName(user.RegisterName)
	if !ok {
		return false
	}

	// hash password
	if err := bcrypt.CompareHashAndPassword([]byte(credential.HashedPassword), []byte(user.Password)); err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "service",
			"target": "user",
			"action": "log in",
		}).Error(err)
		return false
	}
	user.ID = userInDB.ID
	return true
}

func (s *userService) Info(id int) (*model.User, bool) {
	user, _, ok := s.userRepo.Find(id)
	if !ok {
		return nil, false
	}
	return user, true
}

func (s *userService) UpdateInfo(id int, new *model.User) bool {
	_, _, ok := s.userRepo.Find(id)
	if !ok {
		return false
	}
	return s.userRepo.UpdateInfo(id, new)
}

func (s *userService) UpdatePassword(id int, new *model.User) bool {
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(new.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "service",
			"target": "user",
			"action": "update password",
		}).Error(err)
		return false
	}

	if ok := s.userRepo.UpdatePassword(id, &model.Credential{
		HashedPassword: string(hashedPassword),
	}); !ok {
		return false
	}
	return true
}
