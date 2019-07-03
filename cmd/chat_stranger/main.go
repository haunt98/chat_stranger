package main

import (
	"github.com/1612180/chat_stranger/internal/storage"
	"github.com/sirupsen/logrus"

	"github.com/1612180/chat_stranger/pkg/configutils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

const (
	DbDialect = "db.dialect"
	DbUrl     = "db.url"
	Port      = "port"
	GinMode   = "gin.mode"
)

func main() {
	// Load config
	configutils.LoadConfiguration("chat_stranger", "config", "configs")

	// Load database
	db, err := gorm.Open(viper.GetString(DbDialect), viper.GetString(DbUrl))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"event":   "database",
			"dialect": viper.GetString(DbDialect),
			"url":     viper.GetString(DbUrl),
		}).Error(err)
		logrus.Error("Failed to connect to database")
		return
	}

	logrus.WithFields(logrus.Fields{
		"event":   "database",
		"dialect": viper.GetString(DbDialect),
		"url":     viper.GetString(DbUrl),
	}).Info("Connect to database OK")

	defer func() {
		if err := db.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"event": "database",
			}).Error(err)
			logrus.Error("Failed to disconnect to database")
			return
		}

		logrus.WithFields(logrus.Fields{
			"event": "database",
		}).Info("Disconnect to database OK")
	}()

	// Migrate
	storage.MigrateAll(db)

	// Load gin config
	gin.SetMode(viper.GetString(GinMode))

	// Create gin router
	router := gin.Default()

	// Start gin router
	if err := router.Run(":" + viper.GetString(Port)); err != nil {
		logrus.WithFields(logrus.Fields{
			"event": "gin",
			"port":  viper.GetString(Port),
		}).Error(err)
		logrus.Error("Failed to start gin router")
	}
}
