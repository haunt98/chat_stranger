package repository

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/jinzhu/gorm"
)

//go:generate mockgen -destination=../mock/mock_repository/mock_message.go -source=message.go

type MessageRepo interface {
	FetchByTime(roomID int, fromTime time.Time) ([]*model.Message, bool)
	Create(message *model.Message) bool
	Delete(roomID int) bool
}

func NewMessageRepo(db *gorm.DB) MessageRepo {
	return &MessageGorm{db: db}
}

// implement

type MessageGorm struct {
	db *gorm.DB
}

func (g *MessageGorm) FetchByTime(roomID int, fromTime time.Time) ([]*model.Message, bool) {
	var messages []*model.Message
	if err := g.db.Where("room_id = ? AND created_at > ?", roomID, fromTime).
		Find(&messages).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "message",
			"action": "messages",
		}).Error(err)
		return nil, false
	}

	// fill user full name
	for _, message := range messages {
		var user model.User
		if err := g.db.Where("id = ?", message.UserID).
			First(&user).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"event":  "repo",
				"target": "message",
				"action": "messages",
			}).Error(err)
			return nil, false
		}
		message.UserFullName = user.FullName
	}
	return messages, true
}

func (g *MessageGorm) Create(message *model.Message) bool {
	if err := g.db.Create(message).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "message",
			"action": "save",
		}).Error(err)
		return false
	}
	return true
}

func (g *MessageGorm) Delete(roomID int) bool {
	if err := g.db.Where("room_id = ?", roomID).
		Delete(model.Message{}).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "message",
			"action": "delete",
		}).Error(err)
		return false
	}
	return true
}
