package repository

import (
	"github.com/1612180/chat_stranger/models"
)

type IUserRepo interface {
	FetchAll() ([]*models.User, error)
	FindByID(uint) (*models.User, error)
	Create(*models.User) error
	Update(*models.User, *models.User) error
	Delete(*models.User) error
}

type ICredentialRepo interface {
	FetchAll() ([]*models.Credential, error)
	FindByID(id uint) (*models.Credential, error)
	FindByName(string) (*models.Credential, error)
	Create(*models.Credential) error
	Update(*models.Credential, *models.Credential) error
	Delete(*models.Credential) error
}
