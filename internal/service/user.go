package service

import (
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) SignUp(user *model.User) bool {
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

func (s *UserService) LogIn(user *model.User) bool {
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

func (s *UserService) Info(id int) (*model.User, bool) {
	user, _, ok := s.userRepo.Find(id)
	if !ok {
		return nil, false
	}
	return user, true
}

func (s *UserService) UpdateInfo(id int, user *model.User) bool {
	_, _, ok := s.userRepo.Find(id)
	if !ok {
		return false
	}
	return s.userRepo.UpdateInfo(id, user)
}

func (s *UserService) UpdatePassword(userID int, user *model.User) bool {
	_, credential, ok := s.userRepo.Find(userID)
	if !ok {
		return false
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "service",
			"target": "user",
			"action": "update password",
		}).Error(err)
		return false
	}

	if ok := s.userRepo.UpdatePassword(credential.ID, &model.Credential{
		HashedPassword: string(hashedPassword),
	}); !ok {
		return false
	}
	return true
}
