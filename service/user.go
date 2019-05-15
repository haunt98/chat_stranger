package service

import (
	"fmt"

	"github.com/1612180/chat_stranger/models"
	"github.com/1612180/chat_stranger/repository"
)

type UserService struct {
	repo repository.IUserRepo
}

func NewUserService(repo repository.IUserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (userService *UserService) FetchAll() ([]*models.User, error) {
	return userService.repo.FetchAll()
}

func (userService *UserService) FindByID(id uint) (*models.User, error) {
	return userService.repo.FindByID(id)
}

func (userService *UserService) FindByUsername(username string) (*models.User, error) {
	return userService.repo.FindByUsername(username)
}

func (userService *UserService) Create(up *models.UserPOST) error {
	// Username already exist
	_, err := userService.FindByUsername(up.Username)
	if err == nil {
		return fmt.Errorf("Username already exists")
	}

	u, err := up.NewUser()
	if err != nil {
		return err
	}
	return userService.repo.Create(u)
}

func (userService *UserService) UpdateByID(id uint, up *models.UserPOST) error {
	// ID not exist
	uOld, err := userService.FindByID(id)
	if err != nil {
		return fmt.Errorf("ID not exists")
	}

	uNew, err := up.NewUser()
	if err != nil {
		return err
	}

	return userService.repo.Update(uOld, uNew)
}

func (userService *UserService) DeleteByID(id uint) error {
	// ID not exist
	u, err := userService.FindByID(id)
	if err != nil {
		return fmt.Errorf("ID not exists")
	}

	return userService.repo.Delete(u)
}
