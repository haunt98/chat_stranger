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

func (userService *UserService) FetchAll() ([]*models.User, []error) {
	return userService.userRepo.FetchAll()
}

func (userService *UserService) FindByID(id uint) (*models.User, []error) {
	return userService.userRepo.FindByID(id)
}

func (userService *UserService) Create(userUpload *models.UserUpload) (bool, []error) {
	_, errs := userService.credentialRepo.FindByName(userUpload.Authentication.Name)
	if len(errs) == 0 {
		var errs []error
		errs = append(errs, fmt.Errorf("Username already exists"))
		return false, errs
	}

	credential, err := userUpload.Authentication.NewCredential()
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return false, errs
	}
	user := userUpload.NewUser()
	user.Credential = *credential

	return userService.userRepo.Create(user)
}

func (userService *UserService) UpdateInfoByID(id uint, userUpload *models.UserUpload) (bool, []error) {
	userOld, errs := userService.FindByID(id)
	if len(errs) != 0 {
		var errs []error
		errs = append(errs, fmt.Errorf("ID not exists"))
		return false, errs
	}

	userNew := userUpload.NewUser()

	return userService.userRepo.UpdateInfo(userOld, userNew)
}

func (userService *UserService) UpdatePasswordByID(id uint, userUpload *models.UserUpload) (bool, []error) {
	userOld, errs := userService.FindByID(id)
	if len(errs) != 0 {
		var errs []error
		errs = append(errs, fmt.Errorf("ID not exists"))
		return false, errs
	}

	credentialOld, errs := userService.credentialRepo.FindByUser(userOld)
	if len(errs) != 0 {
		return false, errs
	}

	credentialNew, err := userUpload.Authentication.NewCredential()
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return false, errs
	}

	return userService.credentialRepo.UpdatePassword(credentialOld, credentialNew)
}

func (userService *UserService) DeleteByID(id uint) (bool, []error) {
	user, errs := userService.FindByID(id)
	if len(errs) != 0 {
		var errs []error
		errs = append(errs, fmt.Errorf("ID not exists"))
		return false, errs
	}

	status, errs := userService.userRepo.Delete(user)
	if status == false {
		return false, errs
	}

	credential, errs := userService.credentialRepo.FindByUser(user)
	if len(errs) != 0 {
		return false, errs
	}

	status, errs = userService.credentialRepo.Delete(credential)
	if status == false {
		return false, errs
	}

	return true, nil
}
