package repository

import (
	"fmt"
	"testing"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/configwrap"
	"github.com/1612180/chat_stranger/internal/pkg/variable"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestUserGorm_Find(t *testing.T) {
	config := configwrap.NewConfig(variable.TestMode)

	db, err := gorm.Open(config.Get(variable.DbDialect), config.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	userGorm := &userGorm{db: db}

	// create data
	migrate(db, t)

	if err := db.Create(&model.Credential{}).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.User{CredentialID: 1}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		id             int
		wantUser       *model.User
		wantCredential *model.Credential
		wantOK         bool
	}{
		{
			id:             0,
			wantUser:       nil,
			wantCredential: nil,
			wantOK:         false,
		},
		{
			id: 1,
			wantUser: &model.User{
				ID:           1,
				CredentialID: 1,
			},
			wantCredential: &model.Credential{
				ID: 1,
			},
			wantOK: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("id=%d", tc.id), func(t *testing.T) {
			user, credential, ok := userGorm.Find(tc.id)
			assert.Equal(t, tc.wantOK, ok)
			assert.Equal(t, tc.wantUser, user)
			assert.Equal(t, tc.wantCredential, credential)
		})
	}
}

func TestUserGorm_FindByRegisterName(t *testing.T) {
	config := configwrap.NewConfig(variable.TestMode)

	db, err := gorm.Open(config.Get(variable.DbDialect), config.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	userGorm := &userGorm{db: db}

	// create data
	migrate(db, t)

	if err := db.Create(&model.Credential{RegisterName: "a"}).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.User{CredentialID: 1}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		n              string
		wantUser       *model.User
		wantCredential *model.Credential
		wantOK         bool
	}{
		{
			n:              "",
			wantUser:       nil,
			wantCredential: nil,
			wantOK:         false,
		},
		{
			n: "a",
			wantUser: &model.User{
				ID:           1,
				CredentialID: 1,
			},
			wantCredential: &model.Credential{
				ID:           1,
				RegisterName: "a",
			},
			wantOK: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("name=%s", tc.n), func(t *testing.T) {
			user, credential, ok := userGorm.FindByRegisterName(tc.n)
			assert.Equal(t, tc.wantOK, ok)
			assert.Equal(t, tc.wantUser, user)
			assert.Equal(t, tc.wantCredential, credential)
		})
	}
}

func TestUserGorm_Create(t *testing.T) {
	config := configwrap.NewConfig(variable.TestMode)

	db, err := gorm.Open(config.Get(variable.DbDialect), config.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	userGorm := &userGorm{db: db}

	// create data
	migrate(db, t)

	testCases := []struct {
		user       *model.User
		credential *model.Credential
		wantOK     bool
	}{
		{
			user:       nil,
			credential: nil,
			wantOK:     false,
		},
		{
			user:       &model.User{},
			credential: &model.Credential{},
			wantOK:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("user=%v credential=%v", tc.user, tc.credential), func(t *testing.T) {
			ok := userGorm.Create(tc.user, tc.credential)
			assert.Equal(t, tc.wantOK, ok)
		})
	}
}

func TestUserGorm_Delete(t *testing.T) {
	config := configwrap.NewConfig(variable.TestMode)

	db, err := gorm.Open(config.Get(variable.DbDialect), config.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	userGorm := &userGorm{db: db}

	// create data
	migrate(db, t)

	if err := db.Create(&model.Credential{}).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.User{CredentialID: 1}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		id     int
		wantOk bool
	}{
		{
			id:     0,
			wantOk: false,
		},
		{
			id:     1,
			wantOk: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("id=%d", tc.id), func(t *testing.T) {
			ok := userGorm.Delete(tc.id)
			assert.Equal(t, tc.wantOk, ok)
		})
	}
}

func TestUserGorm_UpdateInfo(t *testing.T) {
	config := configwrap.NewConfig(variable.TestMode)

	db, err := gorm.Open(config.Get(variable.DbDialect), config.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	userGorm := &userGorm{db: db}

	// create data
	migrate(db, t)

	if err := db.Create(&model.User{}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		id     int
		user   *model.User
		wantOK bool
	}{
		{
			id:     0,
			user:   nil,
			wantOK: false,
		},
		{
			id:     1,
			user:   &model.User{},
			wantOK: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("id=%d user=%v", tc.id, tc.user), func(t *testing.T) {
			ok := userGorm.UpdateInfo(tc.id, tc.user)
			assert.Equal(t, tc.wantOK, ok)
		})
	}
}

func TestUserGorm_UpdatePassword(t *testing.T) {
	config := configwrap.NewConfig(variable.TestMode)

	db, err := gorm.Open(config.Get(variable.DbDialect), config.Get(variable.DbUrl))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	userGorm := &userGorm{db: db}

	// create data
	migrate(db, t)

	if err := db.Create(&model.Credential{}).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.User{CredentialID: 1}).Error; err != nil {
		t.Error(err)
	}

	testCases := []struct {
		id         int
		credential *model.Credential
		wantOK     bool
	}{
		{
			id:         0,
			credential: nil,
			wantOK:     false,
		},
		{
			id:         1,
			credential: &model.Credential{},
			wantOK:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("id=%d credential=%v", tc.id, tc.credential), func(t *testing.T) {
			ok := userGorm.UpdatePassword(tc.id, tc.credential)
			assert.Equal(t, tc.wantOK, ok)
		})
	}
}
