package service

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/1612180/chat_stranger/models"
	"github.com/1612180/chat_stranger/repository"
)

type UserService struct {
	credentialRepo repository.ICredentialRepo
	userRepo       repository.IUserRepo
}

func NewUserService(credentialRepo repository.ICredentialRepo, userRepo repository.IUserRepo) *UserService {
	return &UserService{
		credentialRepo: credentialRepo,
		userRepo:       userRepo,
	}
}
func (userService *UserService) FetchAll() ([]*models.User, []error) {
	return userService.userRepo.FetchAll()
}

func (userService *UserService) Find(id uint) (*models.User, []error) {
	return userService.userRepo.Find(id)
}

func (userService *UserService) Create(userUpload *models.UserUpload) (uint, []error) {
	return userService.userRepo.Create(userUpload)
}

func (userService *UserService) UpdateInfo(id uint, userUpload *models.UserUpload) []error {
	return userService.userRepo.UpdateInfo(id, userUpload)
}

func (userService *UserService) UpdatePassword(id uint, authentication *models.Authentication) []error {
	return userService.userRepo.UpdatePassword(id, authentication)
}

func (userService *UserService) Delete(id uint) []error {
	return userService.userRepo.Delete(id)
}

func (userService *UserService) Authenticate(authentication *models.Authentication) []error {
	credential, errs := userService.credentialRepo.Find(authentication.Name)
	if len(errs) != 0 {
		return errs
	}

	if _, errs = userService.credentialRepo.TryUser(credential); len(errs) != 0 {
		return errs
	}

	if err := bcrypt.CompareHashAndPassword([]byte(credential.HashedPassword), []byte(authentication.Password)); err != nil {
		var errs []error
		errs = append(errs, err)
		return errs
	}

	return nil
}
