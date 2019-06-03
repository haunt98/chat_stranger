package repository

import (
	"github.com/1612180/chat_stranger/internal/models"
)

type UserRepo interface {
	FetchAll() ([]*models.User, []error)
	Find(int) (*models.User, []error)
	Create(*models.UserUpload) (int, []error)
	UpdateInfo(int, *models.UserUpload) []error
	UpdatePassword(int, *models.Authentication) []error
	Delete(int) []error
}

type AdminRepo interface {
	FetchAll() ([]*models.Admin, []error)
	Find(int) (*models.Admin, []error)
	Create(*models.AdminUpload) (int, []error)
	UpdateInfo(int, *models.AdminUpload) []error
	UpdatePassword(int, *models.Authentication) []error
	Delete(int) []error
}

type CredentialRepo interface {
	Find(string) (*models.Credential, []error)
	TryAdmin(*models.Credential) (*models.Admin, []error)
	TryUser(*models.Credential) (*models.User, []error)
}
