package model

import (
	"github.com/1612180/chat_stranger/internal/pkg/variable"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Migrate(db *gorm.DB) {
	if viper.GetString(variable.DbMode) == "debug" {
		if err := db.DropTableIfExists(
			&Credential{},
			&User{},
			&Admin{},
			&Hobby{},
			&Like{},
			&Room{},
			&Member{},
			&Message{},
		).Error; err != nil {
			logrus.Error(err)
			logrus.WithFields(logrus.Fields{
				"event": "database",
				"mode":  "debug",
			}).Error("Failed to drop all tables in debug mode")
		}
	}

	if err := db.AutoMigrate(
		&Credential{},
		&User{},
		&Admin{},
		&Hobby{},
		&Like{},
		&Room{},
		&Member{},
		&Message{},
	).Error; err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event": "database",
		}).Error("Failed to migrate all tables")
	}
}
