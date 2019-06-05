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

func (s *AdminService) FetchAll() ([]*models.AdminDownload, []error) {
	admins, errs := s.adminRepo.FetchAll()
	if len(errs) != 0 {
		return nil, errs
	}

	var downs []*models.AdminDownload
	for _, admin := range admins {
		downs = append(downs, &models.AdminDownload{
			ID:       admin.ID,
			FullName: admin.FullName,
		})
	}

	return downs, nil
}

func (s *AdminService) Find(id int) (*models.AdminDownload, []error) {
	admin, errs := s.adminRepo.Find(id)
	if len(errs) != 0 {
		return nil, errs
	}

	return &models.AdminDownload{
		ID:       admin.ID,
		FullName: admin.FullName,
	}, nil
}

func (s *AdminService) Create(upload *models.AdminUpload) (int, []error) {
	return s.adminRepo.Create(upload)
}

func (s *AdminService) UpdateInfo(id int, upload *models.AdminUpload) []error {
	return s.adminRepo.UpdateInfo(id, upload)
}

func (s *AdminService) UpdatePassword(id int, auth *models.Authentication) []error {
	return s.adminRepo.UpdatePassword(id, auth)
}

func (s *AdminService) Delete(id int) []error {
	return s.adminRepo.Delete(id)
}

func (s *AdminService) Authenticate(auth *models.Authentication) (*models.AdminDownload, []error) {
	cre, errs := s.credentialRepo.Find(auth.RegName)
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

	return &models.AdminDownload{
		ID:       admin.ID,
		FullName: admin.FullName,
	}, nil
}
