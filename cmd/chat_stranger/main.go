package main

import (
	"github.com/1612180/chat_stranger/internal/handler"
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/env"
	"github.com/1612180/chat_stranger/internal/repository"
	"github.com/1612180/chat_stranger/internal/service"
	"github.com/sirupsen/logrus"

	"github.com/1612180/chat_stranger/pkg/configutils"
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
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event":   "database",
			"dialect": viper.GetString(env.DbDialect),
			"url":     viper.GetString(env.DbUrl),
		}).Error("Failed to connect to database")
		return
	}

	defer func() {
		if err := db.Close(); err != nil {
			logrus.Error(err)
			logrus.WithFields(logrus.Fields{
				"event": "database",
			}).Error("Failed to disconnect to database")
			return
		}
	}()

	// Migrate
	model.Migrate(db)

	// Load repository
	userRepo := repository.NewUserRepository(db)
	roomRepo := repository.NewRoomRepository(db)

	// Load service
	userService := service.NewUserService(userRepo)
	roomService := service.NewRoomService(roomRepo)

	// Load handler
	userHandler := handler.NewUserHandler(userService)
	roomHandler := handler.NewRoomHandler(roomService)

	// Create gin router
	router := handler.NewRouter(userHandler, roomHandler)

	// Start gin router
	if err := router.Run(":" + viper.GetString(env.Port)); err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event": "gin",
			"port":  viper.GetString(env.Port),
		}).Error("Failed to start gin router")
	}
}
