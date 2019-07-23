package repository

import (
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

//go:generate $GOPATH/bin/mockgen -destination=../mock/mock_repository/mock_account.go -source=account.go

type AccountRepo interface {
	FindUserCredential(userID int) (model.User, model.Credential, error)
	FindUserCredentialByRegName(regName string) (model.User, model.Credential, error)
	CreateUserCredential(showName, regName, hashedPass string) (model.User, model.Credential, error)
	UpdateUser(userID int, showName, gender string, birthYear int) (model.User, error)
}

func NewAccountRepo(db *gorm.DB) AccountRepo {
	return &defaultAccountRepo{db: db}
}

// implement

type defaultAccountRepo struct {
	db *gorm.DB
}

func (r *defaultAccountRepo) FindUserCredential(userID int) (model.User, model.Credential, error) {
	var user model.User
	var cred model.Credential
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return user, cred, errors.Wrapf(err, "account repo find user credential with userID=%d failed", userID)
	}
	if err := r.db.Where("id = ?", user.CredentialID).First(&cred).Error; err != nil {
		return user, cred, errors.Wrapf(err, "account repo find user credential with userID=%d failed", userID)
	}
	return user, cred, nil
}

func (r *defaultAccountRepo) FindUserCredentialByRegName(regName string) (model.User, model.Credential, error) {
	var user model.User
	var cred model.Credential
	if err := r.db.Where("register_name = ?", regName).First(&cred).Error; err != nil {
		return user, cred, errors.Wrapf(err, "account repo find user credential with regName=%s failed", regName)
	}
	if err := r.db.Where("credential_id = ?", cred.ID).First(&user).Error; err != nil {
		return user, cred, errors.Wrapf(err, "account repo find user credential with regName=%s failed", regName)
	}
	return user, cred, nil
}

func (r *defaultAccountRepo) CreateUserCredential(showName, regName, hashedPass string) (model.User, model.Credential, error) {
	user := model.User{ShowName: showName}
	cred := model.Credential{RegisterName: regName, HashedPassword: hashedPass}
	if err := r.db.Create(&cred).Error; err != nil {
		return user, cred, errors.Wrapf(err, "account repo create user credential failed")
	}
	user.CredentialID = cred.ID
	if err := r.db.Create(&user).Error; err != nil {
		return user, cred, errors.Wrapf(err, "account repo create user credential failed")
	}
	return user, cred, nil
}

func (r *defaultAccountRepo) UpdateUser(userID int, showName, gender string, birthYear int) (model.User, error) {
	var user model.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return user, errors.Wrapf(err, "account repo update user userID=%d failed", userID)
	}
	if err := r.db.Model(&user).Updates(model.User{ShowName: showName, Gender: gender, BirthYear: birthYear}).Error; err != nil {
		return user, errors.Wrapf(err, "account repo update user userID=%d failed", userID)
	}
	return user, nil
}
