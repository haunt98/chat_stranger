package repository

import (
	"testing"
	"time"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestMessageGorm_FetchByTime(t *testing.T) {
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
		&model.User{},
		&model.Message{},
	).Error; err != nil {
		t.Error(err)
	}
	if err := db.AutoMigrate(
		&model.User{},
		&model.Message{},
	).Error; err != nil {
		t.Error(err)
	}

	messageGorm := messageGorm{db: db}

	messages, ok := messageGorm.FetchByTime(1, time.Time{})
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, len(messages))
}

func TestMessageGorm_Create(t *testing.T) {
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
		&model.Message{},
	).Error; err != nil {
		t.Error(err)
	}
	if err := db.AutoMigrate(
		&model.Message{},
	).Error; err != nil {
		t.Error(err)
	}

	messageGorm := messageGorm{db: db}

	ok := messageGorm.Create(&model.Message{})
	assert.Equal(t, true, ok)
}

func TestMessageGorm_Delete(t *testing.T) {
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
		&model.Message{},
	).Error; err != nil {
		t.Error(err)
	}
	if err := db.AutoMigrate(
		&model.Message{},
	).Error; err != nil {
		t.Error(err)
	}

	messageGorm := messageGorm{db: db}

	ok := messageGorm.Delete(1)
	assert.Equal(t, true, ok)
}
