package repository

import (
	"time"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../mock/mock_repository/mock_room.go -source=room.go

type RoomRepository interface {
	Limit() int
	Exist(id int) bool
	IsEmpty(id int) bool
	FindEmpty() (*model.Room, bool)
	FindNext(old int) (*model.Room, bool)
	FindByUserID(userID int) (*model.Room, bool)
	IsUserFree(userID int) bool
	Create() (*model.Room, bool)
	Join(userID, roomID int) bool
	Leave(userID int) bool
	LatestMessage(roomID int, fromTime time.Time) ([]*model.Message, bool)
	CreateMessage(message *model.Message) bool
	DropMessages(roomID int) bool
	CountMember(roomID int) (int, bool)
}

// implement

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomGorm{db: db}
}

type roomGorm struct {
	db *gorm.DB
}

func (g *roomGorm) Exist(id int) bool {
	if err := g.db.Where("id = ?", id).First(&model.Room{}).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "exist",
		}).Error(err)
		return false
	}
	return true
}

func (g *roomGorm) IsEmpty(id int) bool {
	var count int
	if err := g.db.Model(&model.Member{}).Where("room_id = ?", id).Count(&count).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "is empty",
		}).Error(err)
		return false
	}

	if count >= g.Limit() {
		return false
	}
	return true
}

func (g *roomGorm) Limit() int {
	return 2
}

func (g *roomGorm) FindEmpty() (*model.Room, bool) {
	var rooms []*model.Room
	if err := g.db.Find(&rooms).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "find empty",
		}).Error(err)
		return nil, false
	}

	// find empty room
	for _, room := range rooms {
		if g.IsEmpty(room.ID) {
			return room, true
		}
	}
	return nil, false
}

func (g *roomGorm) FindNext(old int) (*model.Room, bool) {
	var rooms []*model.Room
	if err := g.db.Find(&rooms).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "find empty",
		}).Error(err)
		return nil, false
	}

	// find next room
	for _, room := range rooms {
		if g.IsEmpty(room.ID) && room.ID != old {
			return room, true
		}
	}
	return nil, false
}

func (g *roomGorm) FindByUserID(userID int) (*model.Room, bool) {
	// find members
	var members []*model.Member
	if err := g.db.Where("user_id = ?", userID).
		Find(&members).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "find by user id",
		}).Error(err)
		return nil, false
	}

	// make sure user in 1 room
	if len(members) != 1 {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "find by user id",
		}).Error("user not in 1 room")
		return nil, false
	}

	// find room
	var room model.Room
	if err := g.db.Where("id = ?", members[0].RoomID).
		First(&room).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "find by user id",
		}).Error(err)
		return nil, false
	}
	return &room, true
}

func (g *roomGorm) IsUserFree(userID int) bool {
	var members []*model.Member
	if err := g.db.Where("user_id = ?", userID).
		Find(&members).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "is free",
		}).Error(err)
		return false
	}

	if len(members) != 0 {
		return false
	}
	return true
}

func (g *roomGorm) Create() (*model.Room, bool) {
	var room model.Room
	if err := g.db.Create(&room).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "create",
		}).Error(err)
		return nil, false
	}
	return &room, true
}

func (g *roomGorm) Join(userID, roomID int) bool {
	if err := g.db.Create(&model.Member{
		UserID: userID,
		RoomID: roomID,
	}).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "join",
		}).Error(err)
		return false
	}
	return true
}

func (g *roomGorm) Leave(userID int) bool {
	if err := g.db.Where("user_id = ?", userID).
		Delete(model.Member{}).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "leave",
		}).Error(err)
		return false
	}
	return true
}

func (g *roomGorm) LatestMessage(roomID int, fromTime time.Time) ([]*model.Message, bool) {
	var messages []*model.Message
	if err := g.db.Where("room_id = ? AND created_at > ?", roomID, fromTime).
		Find(&messages).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "message",
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
				"target": "room",
				"action": "message",
			}).Error(err)
			return nil, false
		}
		message.UserFullName = user.FullName
	}

	return messages, true
}

func (g *roomGorm) CreateMessage(message *model.Message) bool {
	if err := g.db.Create(message).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "message",
		}).Error(err)
		return false
	}
	return true
}

func (g *roomGorm) DropMessages(roomID int) bool {
	if err := g.db.Where("room_id = ?", roomID).
		Delete(model.Message{}).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "drop messages",
		}).Error(err)
		return false
	}
	return true
}

func (g *roomGorm) CountMember(roomID int) (int, bool) {
	var count int
	if err := g.db.Model(&model.Member{}).
		Where("room_id = ?", roomID).Count(&count).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "room",
			"action": "count member",
		}).Error(err)
		return 0, false
	}
	return count, true
}
