package service

import (
	"github.com/1612180/chat_stranger/internal/dtos"
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

func (s *AdminService) FetchAll() ([]*dtos.AdminResponse, []error) {
	admins, errs := s.adminRepo.FetchAll()
	if len(errs) != 0 {
		return nil, errs
	}

	var adminRess []*dtos.AdminResponse
	for _, admin := range admins {
		adminRess = append(adminRess, admin.ToResponse())
	}

	return adminRess, nil
}

func (s *AdminService) Find(id int) (*dtos.AdminResponse, []error) {
	admin, errs := s.adminRepo.Find(id)
	if len(errs) != 0 {
		return nil, errs
	}

	return admin.ToResponse(), nil
}

func (s *AdminService) Create(adminReq *dtos.AdminRequest) (int, []error) {
	admin, errs := (&models.Admin{}).FromRequest(adminReq)
	if len(errs) != 0 {
		return 0, errs
	}

	return s.adminRepo.Create(admin)
}

func (s *AdminService) UpdateInfo(id int, adminReq *dtos.AdminRequest) []error {
	admin := (&models.Admin{}).UpdateFromRequest(adminReq)
	return s.adminRepo.UpdateInfo(id, admin)
}

func (s *AdminService) Delete(id int) []error {
	return s.adminRepo.Delete(id)
}

func (s *AdminService) Authenticate(credReq *dtos.CredentialRequest) (*dtos.AdminResponse, []error) {
	cre, errs := s.credentialRepo.Find(credReq.RegName)
	if len(errs) != 0 {
		return nil, errs
	}

	admin, errs := s.credentialRepo.TryAdmin(cre)
	if len(errs) != 0 {
		return nil, errs
	}

	if err := bcrypt.CompareHashAndPassword([]byte(cre.HashedPassword), []byte(credReq.Password)); err != nil {
		var errs []error
		errs = append(errs, err)
		return nil, errs
	}

	return admin.ToResponse(), nil
}
