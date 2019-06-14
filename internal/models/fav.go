package models

import (
	"time"

	"github.com/1612180/chat_stranger/internal/dtos"
)

type Favorite struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Users     []*User `gorm:"many2many:user_favorites;"`
}

func (fav *Favorite) FromRequest(req dtos.FavoriteRequest) *Favorite {
	fav.Name = req.Name

	return fav
}

func (fav *Favorite) ToResponse() *dtos.FavoriteResponse {
	return &dtos.FavoriteResponse{
		Name: fav.Name,
	}
}
