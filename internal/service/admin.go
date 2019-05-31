package service

import (
	"github.com/1612180/chat_stranger/internal/models"
	"github.com/1612180/chat_stranger/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	credentialRepo repository.CredentialRepo
	adminRepo      repository.AdminRepo
}

func NewAdminService(creRepo repository.CredentialRepo, adminRepo repository.AdminRepo) *AdminService {
	return &AdminService{
		credentialRepo: creRepo,
		adminRepo:      adminRepo,
	}
}

func (s *AdminService) FetchAll() ([]*models.Admin, []error) {
	return s.adminRepo.FetchAll()
}

func (s *AdminService) Find(id uint) (*models.Admin, []error) {
	return s.adminRepo.Find(id)
}

func (s *AdminService) Create(upload *models.AdminUpload) (uint, []error) {
	return s.adminRepo.Create(upload)
}

func (s *AdminService) UpdateInfo(id uint, upload *models.AdminUpload) []error {
	return s.adminRepo.UpdateInfo(id, upload)
}

func (s *AdminService) UpdatePassword(id uint, auth *models.Authentication) []error {
	return s.adminRepo.UpdatePassword(id, auth)
}

func (s *AdminService) Delete(id uint) []error {
	return s.adminRepo.Delete(id)
}

func (s *AdminService) Authenticate(auth *models.Authentication) (*models.Admin, []error) {
	cre, errs := s.credentialRepo.Find(auth.Name)
	if len(errs) != 0 {
		return nil, errs
	}

	admin, errs := s.credentialRepo.TryAdmin(cre)
	if len(errs) != 0 {
		return nil, errs
	}

	if err := bcrypt.CompareHashAndPassword([]byte(cre.HashedPassword), []byte(auth.Password)); err != nil {
		var errs []error
		errs = append(errs, err)
		return nil, errs
	}

	return admin, nil
}
