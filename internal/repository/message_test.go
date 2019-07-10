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

	t.Run("message 0", func(t *testing.T) {
		messages, ok := messageGorm.FetchByTime(1, time.Time{})
		assert.Equal(t, true, ok)
		assert.Equal(t, 0, len(messages))
	})

	t.Run("message 1", func(t *testing.T) {
		if err := db.Create(&model.User{
			FullName: "a",
		}).Error; err != nil {
			t.Error(err)
		}

		if err := db.Create(&model.Message{
			Body:   "b",
			RoomID: 1,
			UserID: 1,
		}).Error; err != nil {
			t.Error(err)
		}

		messages, ok := messageGorm.FetchByTime(1, time.Time{})
		assert.Equal(t, true, ok)
		assert.Equal(t, 1, len(messages))
		assert.Equal(t, "a", messages[0].UserFullName)
		assert.Equal(t, "b", messages[0].Body)
	})
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

	migrate(db, t)

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

	migrate(db, t)

	messageGorm := messageGorm{db: db}

	ok := messageGorm.Delete(1)
	assert.Equal(t, true, ok)
}
