package repository

import (
	"fmt"
	"testing"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/config"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestNewAccountRepo(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	assert.Equal(t, &defaultAccountRepo{db: db}, NewAccountRepo(db))
}

func TestDefaultAccountRepo_FindUserCredential(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	accountRepo := defaultAccountRepo{db: db}

	migrate(db, t)
	tempCred := model.Credential{RegisterName: "a", HashedPassword: "b"}
	if err := db.Create(&tempCred).Error; err != nil {
		t.Error(err)
	}
	tempUser := model.User{ShowName: "c", CredentialID: tempCred.ID}
	if err := db.Create(&tempUser).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		inUserID      int
		outUser       model.User
		outCredential model.Credential
		ok            bool
	}{
		{
			inUserID:      tempUser.ID,
			outUser:       tempUser,
			outCredential: tempCred,
			ok:            true,
		},
		{
			inUserID: 0,
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			user, cred, err := accountRepo.FindUserCredential(tc.inUserID)
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, tc.outUser, user)
				assert.Equal(t, tc.outCredential, cred)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDefaultAccountRepo_FindUserCredentialByRegName(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	accountRepo := defaultAccountRepo{db: db}

	migrate(db, t)
	tempCred := model.Credential{RegisterName: "a", HashedPassword: "b"}
	if err := db.Create(&tempCred).Error; err != nil {
		t.Error(err)
	}
	tempUser := model.User{ShowName: "c", CredentialID: tempCred.ID}
	if err := db.Create(&tempUser).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		inRegName     string
		outUser       model.User
		outCredential model.Credential
		ok            bool
	}{
		{
			inRegName:     tempCred.RegisterName,
			outUser:       tempUser,
			outCredential: tempCred,
			ok:            true,
		},
		{
			inRegName: "",
			ok:        false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			user, cred, err := accountRepo.FindUserCredentialByRegName(tc.inRegName)
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, tc.outUser, user)
				assert.Equal(t, tc.outCredential, cred)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDefaultAccountRepo_CreateUserCredential(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	accountRepo := defaultAccountRepo{db: db}

	migrate(db, t)
	if err := db.Create(&model.Credential{RegisterName: "a"}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		inShowName    string
		inRegName     string
		inHashedPass  string
		outUser       model.User
		outCredential model.Credential
		ok            bool
	}{
		{
			inShowName:    "a",
			inRegName:     "b",
			inHashedPass:  "c",
			outUser:       model.User{ID: 1, CredentialID: 2, ShowName: "a"},
			outCredential: model.Credential{ID: 2, RegisterName: "b", HashedPassword: "c"},
			ok:            true,
		},
		{
			inRegName: "a",
			ok:        false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			user, cred, err := accountRepo.CreateUserCredential(tc.inShowName, tc.inRegName, tc.inHashedPass)
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, tc.outUser, user)
				assert.Equal(t, tc.outCredential, cred)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDefaultAccountRepo_UpdateUser(t *testing.T) {
	cfg := config.NewConfig(variable.Test)
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	accountRepo := defaultAccountRepo{db: db}

	migrate(db, t)
	tempCred := model.Credential{RegisterName: "a", HashedPassword: "b"}
	if err := db.Create(&tempCred).Error; err != nil {
		t.Error(err)
	}
	tempUser := model.User{CredentialID: tempCred.ID}
	if err := db.Create(&tempUser).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		inUserID    int
		inShowName  string
		inGender    string
		inBirthYear int
		outUser     model.User
		ok          bool
	}{
		{
			inUserID:    tempUser.ID,
			inShowName:  "c",
			inGender:    "d",
			inBirthYear: 1,
			outUser:     model.User{ID: tempUser.ID, ShowName: "c", Gender: "d", BirthYear: 1, CredentialID: tempCred.ID},
			ok:          true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			user, err := accountRepo.UpdateUser(tc.inUserID, tc.inShowName, tc.inGender, tc.inBirthYear)
			if tc.ok {
				assert.Nil(t, err)
				assert.Equal(t, tc.outUser, user)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
