package repository

import (
	"github.com/1612180/chat_stranger/models"
)

type IUserRepo interface {
	FetchAll() ([]*models.User, []error)
	FindByID(uint) (*models.User, []error)
	Create(*models.User) (bool, []error)
	UpdateInfo(*models.User, *models.User) (bool, []error)
	Delete(*models.User) (bool, []error)
}

type ICredentialRepo interface {
	FetchAll() ([]*models.Credential, []error)
	FindByID(id uint) (*models.Credential, []error)
	FindByName(string) (*models.Credential, []error)
	FindByUser(*models.User) (*models.Credential, []error)
	Create(*models.Credential) (bool, []error)
	UpdatePassword(*models.Credential, *models.Credential) (bool, []error)
	Delete(*models.Credential) (bool, []error)
}
