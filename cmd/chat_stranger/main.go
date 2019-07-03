package main

import (
	"github.com/1612180/chat_stranger/internal/pkg/env"
	"github.com/1612180/chat_stranger/pkg/ginrus"
	"github.com/sirupsen/logrus"

	"github.com/1612180/chat_stranger/pkg/configutils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func main() {
	// Load config
	configutils.LoadConfiguration("chat_stranger", "config", "configs")

	// Load database
	db, err := gorm.Open(viper.GetString(env.DbDialect), viper.GetString(env.DbUrl))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"event":   "database",
			"dialect": viper.GetString(env.DbDialect),
			"url":     viper.GetString(env.DbUrl),
		}).Error(err)
		logrus.Error("Failed to connect to database")
		return
	}

	defer func() {
		if err := db.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"event": "database",
			}).Error(err)
			logrus.Error("Failed to disconnect to database")
			return
		}
	}()

	// Migrate
	// repository.MigrateAll(db)

	// Load gin config
	gin.SetMode(viper.GetString(env.GinMode))

	// Create gin router
	router := gin.New()
	router.Use(ginrus.Logger(), gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello")
	})

	// Start gin router
	if err := router.Run(":" + viper.GetString(env.Port)); err != nil {
		logrus.WithFields(logrus.Fields{
			"event": "gin",
			"port":  viper.GetString(env.Port),
		}).Error(err)
		logrus.Error("Failed to start gin router")
	}
}
