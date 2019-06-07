package service

import (
	"github.com/1612180/chat_stranger/internal/dtos"
	"github.com/1612180/chat_stranger/internal/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/1612180/chat_stranger/internal/repository"
)

type UserService struct {
	credentialRepo repository.CredentialRepo
	userRepo       repository.UserRepo
}

func NewUserService(creRepo repository.CredentialRepo, userRepo repository.UserRepo) *UserService {
	return &UserService{
		credentialRepo: creRepo,
		userRepo:       userRepo,
	}
}
func (s *UserService) FetchAll() ([]*dtos.UserResponse, []error) {
	users, errs := s.userRepo.FetchAll()
	if len(errs) != 0 {
		return nil, errs
	}

	var userRess []*dtos.UserResponse
	for _, user := range users {
		userRess = append(userRess, user.ToResponse())
	}

	return userRess, nil
}

func (s *UserService) Find(id int) (*dtos.UserResponse, []error) {
	user, errs := s.userRepo.Find(id)
	if len(errs) != 0 {
		return nil, errs
	}

	return user.ToResponse(), nil
}

func (s *UserService) Create(userReq *dtos.UserRequest) (int, []error) {
	user, errs := (&models.User{}).FromRequest(userReq)
	if len(errs) != 0 {
		return 0, nil
	}

	return s.userRepo.Create(user)
}

func (s *UserService) UpdateInfo(id int, userReq *dtos.UserRequest) []error {
	user := (&models.User{}).UpdateFromRequest(userReq)
	return s.userRepo.UpdateInfo(id, user)
}

func (s *UserService) Delete(id int) []error {
	return s.userRepo.Delete(id)
}

func (s *UserService) Authenticate(credReq *dtos.CredentialRequest) (*dtos.UserResponse, []error) {
	cre, errs := s.credentialRepo.Find(credReq.RegName)
	if len(errs) != 0 {
		return nil, errs
	}

	user, errs := s.credentialRepo.TryUser(cre)
	if len(errs) != 0 {
		return nil, errs
	}

	if err := bcrypt.CompareHashAndPassword([]byte(cre.HashedPassword), []byte(credReq.Password)); err != nil {
		var errs []error
		errs = append(errs, err)
		return nil, errs
	}

	return user.ToResponse(), nil
}
