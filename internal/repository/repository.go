package repository

import (
	"github.com/1612180/chat_stranger/internal/models"
)

type UserRepo interface {
	FetchAll() ([]*models.User, []error)
	Find(id int) (*models.User, []error)
	Create(*models.User) (int, []error)
	UpdateInfo(int, *models.User) []error
	Delete(id int) []error
}

type AdminRepo interface {
	FetchAll() ([]*models.Admin, []error)
	Find(id int) (*models.Admin, []error)
	Create(*models.Admin) (int, []error)
	UpdateInfo(int, *models.Admin) []error
	Delete(id int) []error
}

type CredentialRepo interface {
	Find(name string) (*models.Credential, []error)
	TryAdmin(*models.Credential) (*models.Admin, []error)
	TryUser(*models.Credential) (*models.User, []error)
}

type FavoriteRepo interface {
	FetchAll() ([]*models.Favorite, []error)
	Find(name string) (*models.Favorite, []error)
	Create(*models.Favorite) (int, []error)
	Delete(id int) []error
}

type RoomRepo interface {
	Limit() int
	FetchAll() ([]*models.Room, []error)
	Find(id int) (*models.Room, []error)
	Create() (int, []error)
	Delete(id int) []error
	FindEmpty() (int, []error)
	NextEmpty(userid, oldroomid int) (int, []error)
	Join(userid, roomid int) []error
	Leave(userid, roomid int) []error
	Check(userid, roomid int) []error
}

type MessageRepo interface {
	FetchAll(roomid int) ([]*models.Message, []error)
	FetchLatest(roomid, latest int) (*models.Message, int, []error)
	Create(*models.Message) []error
}
