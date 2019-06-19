package repository

import (
	"github.com/1612180/chat_stranger/internal/models"
)

type UserRepo interface {
	FetchAll() ([]*models.User, []error)
	Find(int) (*models.User, []error)
	Create(user *models.User) (int, []error)
	UpdateInfo(int, *models.User) []error
	Delete(int) []error
}

type AdminRepo interface {
	FetchAll() ([]*models.Admin, []error)
	Find(int) (*models.Admin, []error)
	Create(*models.Admin) (int, []error)
	UpdateInfo(int, *models.Admin) []error
	Delete(int) []error
}

type CredentialRepo interface {
	Find(string) (*models.Credential, []error)
	TryAdmin(*models.Credential) (*models.Admin, []error)
	TryUser(*models.Credential) (*models.User, []error)
}

type FavoriteRepo interface {
	FetchAll() ([]*models.Favorite, []error)
	Find(string) (*models.Favorite, []error)
	Create(*models.Favorite) (int, []error)
	Delete(int) []error
}

type RoomRepo interface {
	Limit() int
	FetchAll() ([]*models.Room, []error)
	Find(int) (*models.Room, []error)
	Create() (int, []error)
	Delete(int) []error
	FindEmpty() (int, []error)
	Join(int, int) []error
	Leave(int, int) []error
	Check(int, int) []error
}

type MessageRepo interface {
	FetchAll(int) ([]*models.Message, []error)
	FetchLatest(int, int) (*models.Message, int, []error)
	Create(*models.Message) []error
}
