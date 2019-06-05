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
func (s *UserService) FetchAll() ([]*models.UserDownload, []error) {
	users, errs := s.userRepo.FetchAll()
	if len(errs) != 0 {
		return nil, errs
	}

	var downs []*models.UserDownload
	for _, user := range users {
		downs = append(downs, &models.UserDownload{
			ID:        user.ID,
			FullName:  user.FullName,
			Gender:    user.Gender,
			BirthYear: user.BirthYear,
			Introduce: user.Introduce,
		})
	}

	return downs, nil
}

func (s *UserService) Find(id int) (*models.UserDownload, []error) {
	user, errs := s.userRepo.Find(id)
	if len(errs) != 0 {
		return nil, errs
	}

	return &models.UserDownload{
		ID:        user.ID,
		FullName:  user.FullName,
		Gender:    user.Gender,
		BirthYear: user.BirthYear,
		Introduce: user.Introduce,
	}, nil
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

func (s *UserService) Authenticate(auth *models.Authentication) (*models.UserDownload, []error) {
	cre, errs := s.credentialRepo.Find(auth.RegName)
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

	return &models.UserDownload{
		ID:        user.ID,
		FullName:  user.FullName,
		Gender:    user.Gender,
		BirthYear: user.BirthYear,
		Introduce: user.Introduce,
	}, nil
}
