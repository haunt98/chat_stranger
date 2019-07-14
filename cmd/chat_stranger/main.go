package main

import (
	"github.com/1612180/chat_stranger/internal/handler"
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/1612180/chat_stranger/internal/repository"
	"github.com/1612180/chat_stranger/internal/service"
	"github.com/1612180/chat_stranger/pkg/viperwrap"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Load config
	viperwrap.Load("chat_stranger", "config", "configs")

	// Load database
	db, err := gorm.Open(viper.GetString(variable.DbDialect), viper.GetString(variable.DbUrl))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"event":   "database",
			"dialect": viper.GetString(variable.DbDialect),
			"url":     viper.GetString(variable.DbUrl),
		}).Error(err)
		return
	}

	defer func() {
		if err := db.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"event": "database",
			}).Error(err)
			return
		}
	}()

	// Migrate
	model.Migrate(db)

	// Load repository
	userRepo := repository.NewUserRepository(db)
	roomRepo := repository.NewRoomRepository(db)
	memberRepo := repository.NewMemberRepo(db)
	messageRepo := repository.NewMessageRepo(db)

	// Load service
	userService := service.NewUserService(userRepo)
	chatService := service.NewChatService(roomRepo, memberRepo, messageRepo)

	// Load handler
	userHandler := handler.NewUserHandler(userService)
	chatHandler := handler.NewChatHandler(chatService)

	// Create gin router
	router := handler.NewRouter(userHandler, chatHandler, false)

	// Start gin router
	if err := router.Run(":" + viper.GetString(variable.Port)); err != nil {
		logrus.WithFields(logrus.Fields{
			"event": "gin",
			"port":  viper.GetString(variable.Port),
		}).Error(err)
	}
}
