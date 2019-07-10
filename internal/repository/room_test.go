package repository

import (
	"testing"

	"github.com/1612180/chat_stranger/internal/model"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestRoomGorm_Exist(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	roomGorm := roomGorm{db: db}

	t.Run("false", func(t *testing.T) {
		migrate(db, t)

		ok := roomGorm.Exist(1)
		assert.Equal(t, false, ok)
	})

	t.Run("true", func(t *testing.T) {
		migrate(db, t)

		if err := db.Create(&model.Room{}).Error; err != nil {
			t.Error(err)
		}

		ok := roomGorm.Exist(1)
		assert.Equal(t, true, ok)
	})
}

func TestRoomGorm_FindEmpty(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	roomGorm := roomGorm{db: db}

	t.Run("false", func(t *testing.T) {
		migrate(db, t)

		_, ok := roomGorm.FindEmpty()
		assert.Equal(t, false, ok)
	})

	t.Run("true", func(t *testing.T) {
		migrate(db, t)

		if err := db.Create(&model.Room{}).Error; err != nil {
			t.Error(err)
		}

		room, ok := roomGorm.FindEmpty()
		assert.Equal(t, true, ok)
		assert.Equal(t, 1, room.ID)
	})
}

func TestRoomGorm_FindNext(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	roomGorm := roomGorm{db: db}

	t.Run("false", func(t *testing.T) {
		migrate(db, t)

		_, ok := roomGorm.FindNext(1)
		assert.Equal(t, false, ok)
	})

	t.Run("false", func(t *testing.T) {
		migrate(db, t)

		if err := db.Create(&model.Room{}).Error; err != nil {
			t.Error(err)
		}

		_, ok := roomGorm.FindNext(1)
		assert.Equal(t, false, ok)
	})

	t.Run("true", func(t *testing.T) {
		migrate(db, t)

		if err := db.Create(&model.Room{}).Error; err != nil {
			t.Error(err)
		}
		if err := db.Create(&model.Room{}).Error; err != nil {
			t.Error(err)
		}

		room, ok := roomGorm.FindNext(1)
		assert.Equal(t, true, ok)
		assert.Equal(t, 2, room.ID)
	})
}
