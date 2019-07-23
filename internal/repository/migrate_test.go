package repository

import (
	"testing"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/jinzhu/gorm"
)

func migrate(db *gorm.DB, t *testing.T) {
	if err := db.DropTableIfExists(&model.Credential{}, &model.User{},
		&model.Room{}, &model.Member{}, &model.Message{}).Error; err != nil {
		t.Error(err)
	}

	if err := db.AutoMigrate(&model.Credential{}, &model.User{},
		&model.Room{}, &model.Member{}, &model.Message{}).Error; err != nil {
		t.Error(err)
	}
}
