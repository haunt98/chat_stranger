package models

import (
	"time"

	"github.com/1612180/chat_stranger/internal/dtos"
)

type Message struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	RoomID    int
	FromUser  int
	Body      string
}

func (msg *Message) FromRequest(req dtos.MessageRequest) *Message {
	msg.RoomID = req.RoomID
	msg.Body = req.Body
}
