package models

import (
	"time"

	"github.com/1612180/chat_stranger/internal/dtos"
)

type Message struct {
	ID         int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	RoomID     int
	FromUserID int
	Body       string
}

func (msg *Message) FromRequest(fromuserid int, req *dtos.MessageRequest) *Message {
	msg.RoomID = req.RoomID
	msg.FromUserID = fromuserid
	msg.Body = req.Body

	return msg
}

func (msg *Message) ToResponse(fullname string) (*dtos.MessageResponse, []error) {

	return &dtos.MessageResponse{
		RoomID:   msg.RoomID,
		FromUser: fullname,
		Body:     msg.Body,
	}, nil
}
