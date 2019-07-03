package model

import "time"

type Message struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	RoomID    int       `json:"room_id"`
	UserID    int       `json:"user_id"`
}
