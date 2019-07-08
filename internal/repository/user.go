package repository

import (
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../mock/mock_repository/mock_user.go -source=user.go

type UserRepository interface {
	Find(id int) (*model.User, *model.Credential, bool)
	FindByRegisterName(n string) (*model.User, *model.Credential, bool)
	Create(user *model.User, credential *model.Credential) bool
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userGorm{db: db}
}

// implement interface

type userGorm struct {
	db *gorm.DB
}

func (g *userGorm) Find(id int) (*model.User, *model.Credential, bool) {
	var user model.User
	if err := g.db.Where("id = ?", id).First(&user).Error; err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "user",
			"id":     id,
		}).Error("Not found user")
		return nil, nil, false
	}

	var credential model.Credential
	if err := g.db.Where("id = ?", user.CredentialID).First(&credential).Error; err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "credential",
			"id":     user.CredentialID,
		}).Error("Not found credential")
		return nil, nil, false
	}
	return &user, &credential, true
}

func (g *userGorm) FindByRegisterName(n string) (*model.User, *model.Credential, bool) {
	var credential model.Credential
	if err := g.db.Where("register_name = ?", n).First(&credential).Error; err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event":         "repo",
			"target":        "credential",
			"register_name": n,
		}).Error("Not found credential")
		return nil, nil, false
	}

	var user model.User
	if err := g.db.Where("credential_id = ?", credential.ID).First(&user).Error; err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event":         "repo",
			"target":        "user",
			"credential_id": credential.ID,
		}).Error("Not found user")
		return nil, nil, false
	}

	return &user, &credential, true
}

func (g *userGorm) Create(user *model.User, credential *model.Credential) bool {
	tx := g.db.Begin()
	if err := tx.Error; err != nil {
		logrus.Error(err)
		tx.Rollback()
		return false
	}

	if err := tx.Create(credential).Error; err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event":      "repo",
			"target":     "credential",
			"credential": credential,
		}).Error("Failed to create credential")
		tx.Rollback()
		return false
	}

	user.CredentialID = credential.ID
	if err := tx.Create(user).Error; err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "user",
			"user":   user,
		}).Error("Failed to create user")
		tx.Rollback()
		return false
	}

	if err := tx.Commit().Error; err != nil {
		logrus.Error(err)
		return false
	}
	return true
}

func (g *userGorm) Delete(id int) bool {
	tx := g.db.Begin()
	if err := tx.Error; err != nil {
		logrus.Error(err)
		tx.Rollback()
		return false
	}

	// find user
	var user model.User
	if err := tx.Where("id = ?", id).First(&user).Error; err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "user",
			"id":     id,
		}).Error("Not found user")
		tx.Rollback()
		return false
	}

	// delete credential
	if err := tx.Where("id = ?", user.CredentialID).Delete(&model.Credential{}).Error; err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "credential",
			"id":     user.CredentialID,
		}).Error("Failed to delete credential")
		tx.Rollback()
		return false
	}

	// delete user
	if err := tx.Delete(&user).Error; err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "user",
			"user":   user,
		}).Error("Failed to delete user")
		tx.Rollback()
		return false
	}

	if err := tx.Commit().Error; err != nil {
		logrus.Error(err)
		return false
	}
	return true
}
