package service

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/1612180/chat_stranger/internal/models"
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
func (s *UserService) FetchAll() ([]*models.User, []error) {
	return s.userRepo.FetchAll()
}

func (s *UserService) Find(id int) (*models.User, []error) {
	return s.userRepo.Find(id)
}

func (s *UserService) Create(upload *models.UserUpload) (int, []error) {
	return s.userRepo.Create(upload)
}

func (s *UserService) UpdateInfo(id int, upload *models.UserUpload) []error {
	return s.userRepo.UpdateInfo(id, upload)
}

func (s *UserService) UpdatePassword(id int, auth *models.Authentication) []error {
	return s.userRepo.UpdatePassword(id, auth)
}

func (s *UserService) Delete(id int) []error {
	return s.userRepo.Delete(id)
}

func (s *UserService) Authenticate(auth *models.Authentication) (*models.User, []error) {
	cre, errs := s.credentialRepo.Find(auth.Name)
	if len(errs) != 0 {
		return nil, errs
	}

	user, errs := s.credentialRepo.TryUser(cre)
	if len(errs) != 0 {
		return nil, errs
	}

	if err := bcrypt.CompareHashAndPassword([]byte(cre.HashedPassword), []byte(auth.Password)); err != nil {
		var errs []error
		errs = append(errs, err)
		return nil, errs
	}

	return user, nil
}
