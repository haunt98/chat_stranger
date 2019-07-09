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

	if err := db.DropTableIfExists(
		&model.Member{},
	).Error; err != nil {
		t.Error(err)
	}
	if err := db.AutoMigrate(
		&model.Member{},
	).Error; err != nil {
		t.Error(err)
	}

	memberGorm := memberGorm{db: db}

	ok := memberGorm.Create(1, 1)
	assert.Equal(t, true, ok)
}
