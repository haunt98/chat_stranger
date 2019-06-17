package models

import (
	"time"

	"github.com/1612180/chat_stranger/internal/dtos"
)

type Room struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Users     []*User `gorm:"many2many:room_users;"`
}

func (room *Room) ToResponse() *dtos.RoomResponse {
	return &dtos.RoomResponse{
		ID: room.ID,
	}
}
