package repository

import (
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../mock/mock_repository/mock_user.go -source=user.go

type UserRepository interface {
	Find(id int) (*model.User, bool)
	Create(user *model.User) bool
	Delete(id int) bool
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userGorm{db: db}
}

// implement interface

type userGorm struct {
	db *gorm.DB
}

func (g *userGorm) Find(id int) (*model.User, bool) {
	var user model.User
	if err := g.db.Where("id = ?", id).First(&user).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "user",
			"id":     id,
		}).Error(err)
		logrus.Error("Not found user")
		return nil, false
	}
	return &user, true
}

func (g *userGorm) Create(user *model.User) bool {
	return false
}

func (g *userGorm) Delete(id int) bool {
	return false
}
