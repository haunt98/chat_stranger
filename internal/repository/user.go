package repository

import (
	"github.com/1612180/chat_stranger/internal/model"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//go:generate $GOPATH/bin/mockgen -destination=../mock/mock_repository/mock_user.go -source=user.go

type UserRepository interface {
	Find(id int) (*model.User, *model.Credential, bool)
	FindByRegisterName(n string) (*model.User, *model.Credential, bool)
	Create(user *model.User, credential *model.Credential) bool
	UpdateInfo(id int, new *model.User) bool
	UpdatePassword(id int, new *model.Credential) bool
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userGorm{db: db}
}

// implement interface

type userGorm struct {
	db *gorm.DB
}

func (g *userGorm) Find(id int) (*model.User, *model.Credential, bool) {
	// find user
	var user model.User
	if err := g.db.Where("id = ?", id).First(&user).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "user",
			"action": "find",
		}).Error(err)
		return nil, nil, false
	}

	// find credential
	var credential model.Credential
	if err := g.db.Where("id = ?", user.CredentialID).First(&credential).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "user",
			"action": "find",
		}).Error(err)
		return nil, nil, false
	}
	return &user, &credential, true
}

func (g *userGorm) FindByRegisterName(n string) (*model.User, *model.Credential, bool) {
	// find credential
	var credential model.Credential
	if err := g.db.Where("register_name = ?", n).First(&credential).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "user",
			"action": "find",
		}).Error(err)
		return nil, nil, false
	}

	// find user
	var user model.User
	if err := g.db.Where("credential_id = ?", credential.ID).First(&user).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "user",
			"action": "find",
		}).Error(err)
		return nil, nil, false
	}
	return &user, &credential, true
}

func (g *userGorm) Create(user *model.User, credential *model.Credential) bool {
	if user == nil || credential == nil {
		return false
	}

	tx := g.db.Begin()
	if err := tx.Error; err != nil {
		logrus.Error(err)
		tx.Rollback()
		return false
	}

	// create credential
	if err := tx.Create(credential).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "user",
			"action": "create",
		}).Error(err)
		tx.Rollback()
		return false
	}

	// create user
	user.CredentialID = credential.ID
	if err := tx.Create(user).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "user",
			"action": "create",
		}).Error(err)
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
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "user",
			"action": "delete",
		}).Error(err)
		tx.Rollback()
		return false
	}

	// delete credential
	if err := tx.Where("id = ?", user.CredentialID).Delete(&model.Credential{}).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "user",
			"action": "delete",
		}).Error(err)
		tx.Rollback()
		return false
	}

	// delete user
	if err := tx.Delete(&user).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "user",
			"action": "delete",
		}).Error(err)
		tx.Rollback()
		return false
	}

	if err := tx.Commit().Error; err != nil {
		logrus.Error(err)
		return false
	}
	return true
}

func (g *userGorm) UpdateInfo(id int, new *model.User) bool {
	if new == nil {
		return false
	}

	var old model.User
	if err := g.db.Where(&model.User{ID: id}).First(&old).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "new",
			"action": "update info",
		}).Error(err)
		return false
	}

	old.FullName = new.FullName
	old.Gender = new.Gender
	old.BirthYear = new.BirthYear

	if err := g.db.Save(&old).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "new",
			"action": "update info",
		}).Error(err)
		return false
	}
	return true
}

func (g *userGorm) UpdatePassword(id int, new *model.Credential) bool {
	if new == nil {
		return false
	}

	var user model.User
	if err := g.db.Where(&model.User{ID: id}).First(&user).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "new",
			"action": "update info",
		}).Error(err)
		return false
	}

	var old model.Credential
	if err := g.db.Where(&model.Credential{ID: user.CredentialID}).First(&old).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "new",
			"action": "update info",
		}).Error(err)
		return false
	}

	old.HashedPassword = new.HashedPassword

	if err := g.db.Save(&old).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "repo",
			"target": "new",
			"action": "update info",
		}).Error(err)
		return false
	}
	return true
}
