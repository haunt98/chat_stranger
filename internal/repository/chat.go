package repository

import (
	"fmt"
	"time"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

//go:generate $GOPATH/bin/mockgen -destination=../mock/mock_repository/mock_chat.go -source=chat.go

type ChatRepo interface {
	CreateRoom() (model.Room, error)
	FindRooms() ([]model.Room, error)
	FindRoomOfUser(userID int) (model.Room, error)
	CountMembersInRoom(roomID int) (int, error)
	FindUsersInRoom(roomID int) ([]model.User, error)
	CountMembersOfUser(userID int) (int, error)
	CreateMember(userID, roomID int) (model.Member, error)
	DeleteMember(userID, roomID int) error
	CreateMessage(roomID, userID int, body string) (model.Message, error)
	DeleteMessagesOfRoom(roomID int) error
	FindMessagesOfRoomFromTime(roomID int, from time.Time) ([]model.Message, error)
}

func NewChatRepo(db *gorm.DB) ChatRepo {
	return &defaultChatRepo{db: db}
}

// implement

type defaultChatRepo struct {
	db *gorm.DB
}

func (r *defaultChatRepo) CreateRoom() (model.Room, error) {
	var room model.Room
	if err := r.db.Create(&room).Error; err != nil {
		return room, errors.Wrap(err, "chat repo: create room failed")
	}
	return room, nil
}

func (r *defaultChatRepo) FindRooms() ([]model.Room, error) {
	var rooms []model.Room
	if err := r.db.Find(&rooms).Error; err != nil {
		return nil, errors.Wrap(err, "chat repo: find rooms failed")
	}
	return rooms, nil
}

func (r *defaultChatRepo) FindRoomOfUser(userID int) (model.Room, error) {
	if err := r.db.Where("id = ?", userID).First(&model.User{}).Error; err != nil {
		return model.Room{}, errors.Wrap(err, "chat repo: find room of user failed")
	}

	var members []model.Member
	if err := r.db.Where("user_id = ?", userID).Find(&members).Error; err != nil {
		return model.Room{}, errors.Wrap(err, "chat repo: find room of user failed")
	}

	if len(members) == 0 {
		return model.Room{}, errors.Wrap(fmt.Errorf("user in no room"), "chat repo: find room of user failed")
	}
	if len(members) > 1 {
		return model.Room{}, errors.Wrap(fmt.Errorf("user in many rooms"), "chat repo: find room of user failed")
	}

	var room model.Room
	if err := r.db.Where("id = ?", members[0].RoomID).First(&room).Error; err != nil {
		return model.Room{}, errors.Wrap(err, "chat repo: find room of user failed")
	}
	return room, nil
}

func (r *defaultChatRepo) CountMembersInRoom(roomID int) (int, error) {
	if err := r.db.Where("id = ?", roomID).First(&model.Room{}).Error; err != nil {
		return 0, errors.Wrap(err, "chat repo: count members in room failed")
	}

	var count int
	if err := r.db.Model(&model.Member{}).Where("room_id = ?", roomID).Count(&count).Error; err != nil {
		return 0, errors.Wrap(err, "chat repo: count members in room failed")
	}
	return count, nil
}

func (r *defaultChatRepo) FindUsersInRoom(roomID int) ([]model.User, error) {
	if err := r.db.Where("id = ?", roomID).First(&model.Room{}).Error; err != nil {
		return nil, errors.Wrap(err, "chat repo: find users in room failed")
	}

	var members []model.Member
	if err := r.db.Where("room_id = ?", roomID).Find(&members).Error; err != nil {
		return nil, errors.Wrap(err, "chat repo: find users in room failed")
	}

	users := make([]model.User, 0, len(members))
	for _, member := range members {
		var user model.User
		if err := r.db.Where("id = ?", member.UserID).First(&user).Error; err != nil {
			return nil, errors.Wrap(err, "chat repo: find users in room failed")
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *defaultChatRepo) CountMembersOfUser(userID int) (int, error) {
	if err := r.db.Where("id = ?", userID).First(&model.User{}).Error; err != nil {
		return 0, errors.Wrap(err, "chat repo: count members of user failed")
	}

	var count int
	if err := r.db.Model(&model.Member{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return 0, errors.Wrap(err, "chat repo: count members of user failed")
	}
	return count, nil
}

func (r *defaultChatRepo) CreateMember(userID, roomID int) (model.Member, error) {
	if err := r.db.Where("id = ?", userID).First(&model.User{}).Error; err != nil {
		return model.Member{}, errors.Wrap(err, "chat repo: create member failed")
	}

	if err := r.db.Where("id = ?", roomID).First(&model.Room{}).Error; err != nil {
		return model.Member{}, errors.Wrap(err, "chat repo: create member failed")
	}

	if !r.db.Where("user_id = ? AND room_id = ?", userID, roomID).First(&model.Member{}).RecordNotFound() {
		return model.Member{}, errors.Wrap(fmt.Errorf("member userID=%d roomID=%d already exist", userID, roomID),
			"chat repo: create member failed")
	}

	member := model.Member{UserID: userID, RoomID: roomID}
	if err := r.db.Create(&member).Error; err != nil {
		return model.Member{}, errors.Wrap(err, "chat repo: create member failed")
	}
	return member, nil
}

func (r *defaultChatRepo) DeleteMember(userID, roomID int) error {
	if err := r.db.Where("id = ?", userID).First(&model.User{}).Error; err != nil {
		return errors.Wrap(err, "chat repo: delete member failed")
	}

	if err := r.db.Where("id = ?", roomID).First(&model.Room{}).Error; err != nil {
		return errors.Wrap(err, "chat repo: delete member failed")
	}

	var member model.Member
	if err := r.db.Where("user_id = ? AND room_id = ?", userID, roomID).First(&member).Error; err != nil {
		return errors.Wrap(err, "chat repo: delete member failed")
	}

	if err := r.db.Delete(&member).Error; err != nil {
		return errors.Wrap(err, "chat repo: delete member failed")
	}
	return nil
}

func (r *defaultChatRepo) CreateMessage(roomID, userID int, body string) (model.Message, error) {
	msg := model.Message{RoomID: roomID, UserID: userID, Body: body}

	if err := r.db.Where("id = ?", msg.RoomID).First(&model.Room{}).Error; err != nil {
		return msg, errors.Wrap(err, "chat repo: create message failed")
	}

	var user model.User
	if err := r.db.Where("id = ?", msg.UserID).First(&user).Error; err != nil {
		return msg, errors.Wrap(err, "chat repo: create message failed")
	}
	msg.UserShowName = user.ShowName

	if err := r.db.Create(&msg).Error; err != nil {
		return msg, errors.Wrap(err, "chat repo: create message failed")
	}
	return msg, nil
}

func (r *defaultChatRepo) DeleteMessagesOfRoom(roomID int) error {
	if err := r.db.Where("id = ?", roomID).First(&model.Room{}).Error; err != nil {
		return errors.Wrap(err, "chat repo: delete messages of room failed")
	}

	if err := r.db.Where("room_id = ?", roomID).Delete(model.Message{}).Error; err != nil {
		return errors.Wrap(err, "chat repo: delete messages of room failed")
	}
	return nil
}

func (r *defaultChatRepo) FindMessagesOfRoomFromTime(roomID int, from time.Time) ([]model.Message, error) {
	if err := r.db.Where("id = ?", roomID).First(&model.Room{}).Error; err != nil {
		return nil, errors.Wrap(err, "chat repo: find messages of room from time failed")
	}

	var msgs []model.Message
	if err := r.db.Where("room_id = ? AND created_at > ?", roomID, from).Find(&msgs).Error; err != nil {
		return nil, errors.Wrap(err, "chat repo: find messages of room from time failed")
	}

	for _, msg := range msgs {
		var user model.User
		if err := r.db.Where("id = ?", msg.UserID).First(&user).Error; err != nil {
			return nil, errors.Wrap(err, "chat repo: find messages of room from time failed")
		}
		msg.UserShowName = user.ShowName
	}
	return msgs, nil
}
