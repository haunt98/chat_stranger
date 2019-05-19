package service

import (
	"fmt"

	"github.com/1612180/chat_stranger/models"
	"github.com/1612180/chat_stranger/repository"
)

type UserService struct {
	credentialRepo  repository.ICredentialRepo
	userRepo        repository.IUserRepo
	transactionRepo repository.ITransactionRepo
}

func NewUserService(credentialRepo repository.ICredentialRepo,
	userRepo repository.IUserRepo,
	transactionRepo repository.ITransactionRepo) *UserService {
	return &UserService{
		credentialRepo:  credentialRepo,
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
	}
}

func (userService *UserService) FetchAll() ([]*models.User, []error) {
	return userService.userRepo.FetchAll()
}

func (userService *UserService) FindByID(id uint) (*models.User, []error) {
	return userService.userRepo.FindByID(id)
}

func (userService *UserService) Create(userUpload *models.UserUpload) []error {
	_, errs := userService.credentialRepo.FindByName(userUpload.Authentication.Name)
	if len(errs) == 0 {
		var errs []error
		errs = append(errs, fmt.Errorf("Username already exists"))
		return errs
	}

	credential, err := userUpload.Authentication.NewCredential()
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return errs
	}
	user := userUpload.NewUser()
	user.Credential = *credential

	return userService.userRepo.Create(user)
}

func (userService *UserService) UpdateInfoByID(id uint, userUpload *models.UserUpload) []error {
	userOld, errs := userService.FindByID(id)
	if len(errs) != 0 {
		return errs
	}

	userNew := userUpload.NewUser()

	return userService.userRepo.UpdateInfo(userOld, userNew)
}

func (userService *UserService) UpdatePasswordByID(id uint, userUpload *models.UserUpload) []error {
	userOld, errs := userService.FindByID(id)
	if len(errs) != 0 {
		return errs
	}

	credentialOld, errs := userService.credentialRepo.FindByUser(userOld)
	if len(errs) != 0 {
		return errs
	}

	credentialNew, err := userUpload.Authentication.NewCredential()
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return errs
	}

	return userService.credentialRepo.UpdatePassword(credentialOld, credentialNew)
}

func (userService *UserService) DeleteByID(id uint) []error {
	return userService.transactionRepo.DeleteUserWithCredentialByUserID(id)
}

func (userService *UserService) Authenticate(authentication *models.Authentication) []error {
	credential, errs := userService.credentialRepo.FindByName(authentication.Name)
	if len(errs) != 0 {
		var errs []error
		errs = append(errs, fmt.Errorf("Name or password is incorrect"))
		return errs
	}

	err := authentication.Authenticate(credential)
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return errs
	}

	return nil
}
