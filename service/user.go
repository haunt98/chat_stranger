package service

import (
	"fmt"

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

func (userService *UserService) FetchAll() ([]*models.User, error) {
	return userService.userRepo.FetchAll()
}

func (userService *UserService) FindByID(id uint) (*models.User, error) {
	return userService.userRepo.FindByID(id)
}

func (userService *UserService) Create(userUpload *models.UserUpload) error {
	_, err := userService.credentialRepo.FindByName(userUpload.Authentication.Name)
	if err == nil {
		return fmt.Errorf("Username already exists")
	}

	u, err := userUpload.NewUser()
	if err != nil {
		return err
	}
	return userService.userRepo.Create(u)
}

func (userService *UserService) UpdateByID(id uint, userUpload *models.UserUpload) error {
	uOld, err := userService.FindByID(id)
	if err != nil {
		return fmt.Errorf("ID not exists")
	}

	uNew, err := userUpload.NewUser()
	if err != nil {
		return err
	}

	return userService.userRepo.Update(uOld, uNew)
}

func (userService *UserService) DeleteByID(id uint) error {
	u, err := userService.FindByID(id)
	if err != nil {
		return fmt.Errorf("ID not exists")
	}

	return userService.userRepo.Delete(u)
}
