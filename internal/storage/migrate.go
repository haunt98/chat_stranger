package storage

import (
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const DbMode = "db.mode"

func MigrateAll(db *gorm.DB) {
	if viper.GetString(DbMode) == "debug" {
		logrus.WithFields(logrus.Fields{
			"event": "databae",
		}).Info("Connect database in debug mode, drop all tables")
		db.DropTableIfExists(&model.Credential{})
		db.DropTableIfExists(&model.Admin{})
		db.DropTableIfExists(&model.User{})

		db.DropTableIfExists(&model.Hobby{})
		db.DropTableIfExists(&model.Like{})

		db.DropTableIfExists(&model.Room{})
		db.DropTableIfExists(&model.Member{})
		db.DropTableIfExists(&model.Message{})
	}

	db.AutoMigrate(&model.Credential{})
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.User{})

	db.AutoMigrate(&model.Hobby{})
	db.AutoMigrate(&model.Like{})

	db.AutoMigrate(&model.Room{})
	db.AutoMigrate(&model.Member{})
	db.AutoMigrate(&model.Message{})

	logrus.WithFields(logrus.Fields{
		"event": "database",
	}).Info("Migrate all tables")
}
