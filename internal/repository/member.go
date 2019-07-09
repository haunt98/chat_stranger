package repository

import (
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../mock/mock_repository/mock_member.go -source=member.go

type MemberRepo interface {
	Create(userID, roomID int) bool
	Delete(userID int) bool
	CountByRoom(roomID int) (int, bool)
	CountByUser(userID int) (int, bool)
}

func NewMemberRepo(db *gorm.DB) MemberRepo {
	return &memberGorm{db: db}
}

// implement

type memberGorm struct {
	db *gorm.DB
}

func (g *memberGorm) Create(userID, roomID int) bool {
	if err := g.db.Create(&model.Member{
		UserID: userID,
		RoomID: roomID,
	}).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "member",
			"action": "save",
		}).Error(err)
		return false
	}
	return true
}

func (g *memberGorm) Delete(userID int) bool {
	if err := g.db.Where("user_id = ?", userID).
		Delete(&model.Member{}).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "member",
			"action": "delete",
		}).Error(err)
		return false
	}
	return true
}

func countByRoom(db *gorm.DB, roomID int) (int, bool) {
	var count int
	if err := db.Model(&model.Member{}).Where("room_id = ?", roomID).
		Count(&count).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"action": "count by room",
		}).Error(err)
		return 0, false
	}
	return count, true
}

func (g *memberGorm) CountByRoom(roomID int) (int, bool) {
	return countByRoom(g.db, roomID)
}

func countByUser(db *gorm.DB, userID int) (int, bool) {
	var count int
	if err := db.Model(&model.Member{}).Where("user_id = ?", userID).
		Count(&count).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "member",
			"action": "count by user",
		}).Error(err)
		return 0, false
	}
	return count, true
}

func (g *memberGorm) CountByUser(userID int) (int, bool) {
	return countByUser(g.db, userID)
}
