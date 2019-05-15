package repository

import (
	"github.com/1612180/chat_stranger/models"
)

type IUserRepo interface {
	FetchAll() ([]*models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Create(*models.User) error
	Update(*models.User, *models.User) error
	Delete(*models.User) error
}
