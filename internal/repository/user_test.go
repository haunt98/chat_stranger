package repository

import (
	"testing"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestUserGorm_Find(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	if err := db.DropTableIfExists(
		&model.Credential{},
		&model.User{},
	).Error; err != nil {
		t.Error(err)
	}
	if err := db.AutoMigrate(
		&model.Credential{},
		&model.User{},
	).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.Credential{}).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.User{CredentialID: 1}).Error; err != nil {
		t.Error(err)
	}

	userGorm := &userGorm{db: db}
	user, credential, ok := userGorm.Find(1)
	assert.Equal(t, true, ok)
	assert.Equal(t, &model.User{ID: 1, CredentialID: 1}, user)
	assert.Equal(t, &model.Credential{ID: 1}, credential)
}

func TestUserGorm_FindByRegisterName(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	if err := db.DropTableIfExists(
		&model.Credential{},
		&model.User{},
	).Error; err != nil {
		t.Error(err)
	}
	if err := db.AutoMigrate(
		&model.Credential{},
		&model.User{},
	).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.Credential{RegisterName: "a"}).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.User{CredentialID: 1}).Error; err != nil {
		t.Error(err)
	}

	userGorm := &userGorm{db: db}
	user, credential, ok := userGorm.FindByRegisterName("a")
	assert.Equal(t, true, ok)
	assert.Equal(t, &model.User{ID: 1, CredentialID: 1}, user)
	assert.Equal(t, &model.Credential{ID: 1, RegisterName: "a"}, credential)

}

func TestUserGorm_Create(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	if err := db.DropTableIfExists(
		&model.Credential{},
		&model.User{},
	).Error; err != nil {
		t.Error(err)
	}
	if err := db.AutoMigrate(
		&model.Credential{},
		&model.User{},
	).Error; err != nil {
		t.Error(err)
	}

	userGorm := &userGorm{db: db}
	var user model.User
	var credential model.Credential
	ok := userGorm.Create(&user, &credential)
	assert.Equal(t, true, ok)
	assert.Equal(t, credential.ID, user.ID)
}

func TestUserGorm_Delete(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	if err := db.DropTableIfExists(
		&model.Credential{},
		&model.User{},
	).Error; err != nil {
		t.Error(err)
	}
	if err := db.AutoMigrate(
		&model.Credential{},
		&model.User{},
	).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.Credential{}).Error; err != nil {
		t.Error(err)
	}
	if err := db.Create(&model.User{CredentialID: 1}).Error; err != nil {
		t.Error(err)
	}

	userGorm := &userGorm{db: db}
	ok := userGorm.Delete(1)
	assert.Equal(t, true, ok)
}

func TestUserGorm_UpdateInfo(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	if err := db.DropTableIfExists(
		&model.Credential{},
		&model.User{},
	).Error; err != nil {
		t.Error(err)
	}
	if err := db.AutoMigrate(
		&model.Credential{},
		&model.User{},
	).Error; err != nil {
		t.Error(err)
	}

	userGorm := &userGorm{db: db}
	ok := userGorm.UpdateInfo(1, &model.User{})
	assert.Equal(t, true, ok)
}

func TestUserGorm_UpdatePassword(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	if err := db.DropTableIfExists(
		&model.Credential{},
		&model.User{},
	).Error; err != nil {
		t.Error(err)
	}
	if err := db.AutoMigrate(
		&model.Credential{},
		&model.User{},
	).Error; err != nil {
		t.Error(err)
	}

	userGorm := &userGorm{db: db}
	ok := userGorm.UpdatePassword(1, &model.Credential{})
	assert.Equal(t, true, ok)
}
