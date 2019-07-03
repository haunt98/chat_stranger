package service

import (
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func (s *UserService) Find(id int) (*model.User, bool) {
	return s.userRepo.Find(id)
}
