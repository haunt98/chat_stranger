package repository

import (
	"github.com/1612180/chat_stranger/models"
)

type IUserRepo interface {
	FetchAll() ([]*models.User, []error)
	FindByID(uint) (*models.User, []error)
	Create(*models.User) []error
	UpdateInfo(*models.User, *models.User) []error
	Delete(*models.User) []error
}

type ICredentialRepo interface {
	FetchAll() ([]*models.Credential, []error)
	FindByID(id uint) (*models.Credential, []error)
	FindByName(string) (*models.Credential, []error)
	FindByUser(*models.User) (*models.Credential, []error)
	Create(*models.Credential) []error
	UpdatePassword(*models.Credential, *models.Credential) []error
	Delete(*models.Credential) []error
}

type ITransactionRepo interface {
	DeleteUserWithCredentialByUserID(id uint) []error
}
