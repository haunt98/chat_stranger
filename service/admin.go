package service

import (
	"github.com/1612180/chat_stranger/models"
	"github.com/1612180/chat_stranger/repository"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	credentialRepo repository.ICredentialRepo
	adminRepo      repository.IAdminRepo
}

func NewAdminService(credentialRepo repository.ICredentialRepo, adminRepo repository.IAdminRepo) *AdminService {
	return &AdminService{
		credentialRepo: credentialRepo,
		adminRepo:      adminRepo,
	}
}

func (adminService *AdminService) FetchAll() ([]*models.Admin, []error) {
	return adminService.adminRepo.FetchAll()
}

func (adminService *AdminService) Find(id uint) (*models.Admin, []error) {
	return adminService.adminRepo.Find(id)
}

func (adminService *AdminService) Create(adminUpload *models.AdminUpload) (uint, []error) {
	return adminService.adminRepo.Create(adminUpload)
}

func (adminService *AdminService) UpdateInfo(id uint, adminUpload *models.AdminUpload) []error {
	return adminService.adminRepo.UpdateInfo(id, adminUpload)
}

func (adminService *AdminService) UpdatePassword(id uint, authentication *models.Authentication) []error {
	return adminService.adminRepo.UpdatePassword(id, authentication)
}

func (adminService *AdminService) Delete(id uint) []error {
	return adminService.adminRepo.Delete(id)
}

func (adminService *AdminService) Authenticate(authentication *models.Authentication) []error {
	credential, errs := adminService.credentialRepo.Find(authentication.Name)
	if len(errs) != 0 {
		return errs
	}

	if _, errs = adminService.credentialRepo.TryAdmin(credential); len(errs) != 0 {
		return errs
	}

	if err := bcrypt.CompareHashAndPassword([]byte(credential.HashedPassword), []byte(authentication.Password)); err != nil {
		var errs []error
		errs = append(errs, err)
		return errs
	}

	return nil
}
