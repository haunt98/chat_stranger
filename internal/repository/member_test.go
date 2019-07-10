package repository

import (
	"testing"

	"github.com/1612180/chat_stranger/internal/model"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestMemberGorm_Create(t *testing.T) {
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

	memberGorm := memberGorm{db: db}

	ok := memberGorm.Create(1, 1)
	assert.Equal(t, true, ok)
}

func TestMemberGorm_Delete(t *testing.T) {
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

	memberGorm := memberGorm{db: db}

	ok := memberGorm.Delete(1)
	assert.Equal(t, true, ok)
}

func TestMemberGorm_CountByRoom(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	memberGorm := memberGorm{db: db}

	t.Run("count 0", func(t *testing.T) {
		migrate(db, t)

		count, ok := memberGorm.CountByRoom(1)
		assert.Equal(t, true, ok)
		assert.Equal(t, 0, count)
	})

	t.Run("count 1", func(t *testing.T) {
		migrate(db, t)

		if err := db.Create(&model.Member{
			UserID: 1,
			RoomID: 1,
		}).Error; err != nil {
			t.Error(err)
		}

		count, ok := memberGorm.CountByRoom(1)
		assert.Equal(t, true, ok)
		assert.Equal(t, 1, count)
	})
}

func TestMemberGorm_CountByUser(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}()

	memberGorm := memberGorm{db: db}

	t.Run("count 0", func(t *testing.T) {
		migrate(db, t)

		count, ok := memberGorm.CountByUser(1)
		assert.Equal(t, true, ok)
		assert.Equal(t, 0, count)
	})

	t.Run("count 1", func(t *testing.T) {
		migrate(db, t)

		if err := db.Create(&model.Member{
			UserID: 1,
			RoomID: 1,
		}).Error; err != nil {
			t.Error(err)
		}

		count, ok := memberGorm.CountByUser(1)
		assert.Equal(t, true, ok)
		assert.Equal(t, 1, count)
	})

}
