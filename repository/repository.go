package repository

import (
	"github.com/1612180/chat_stranger/models"
)

type IUserRepo interface {
	FetchAll() ([]*models.User, []error)
	Find(uint) (*models.User, []error)
	Create(*models.UserUpload) (uint, []error)
	UpdateInfo(uint, *models.UserUpload) []error
	UpdatePassword(uint, *models.Authentication) []error
	Delete(uint) []error
}

type IAdminRepo interface {
	FetchAll() ([]*models.Admin, []error)
	Find(uint) (*models.Admin, []error)
	Create(*models.AdminUpload) (uint, []error)
	UpdateInfo(uint, *models.AdminUpload) []error
	UpdatePassword(uint, *models.Authentication) []error
	Delete(uint) []error
}

type ICredentialRepo interface {
	Find(string) (*models.Credential, []error)
	TryAdmin(*models.Credential) (*models.Admin, []error)
	TryUser(*models.Credential) (*models.User, []error)
}
