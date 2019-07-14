package main

import (
	"github.com/1612180/chat_stranger/internal/handler"
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/configwrap"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/1612180/chat_stranger/internal/pkg/viperwrap"
	"github.com/1612180/chat_stranger/internal/repository"
	"github.com/1612180/chat_stranger/internal/service"
	"github.com/1612180/chat_stranger/pkg/gormrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load config
	viperwrap.Load("chat_stranger", "config", "configs")
	config := configwrap.NewConfig("viper")

	// Load database
	db, err := gorm.Open(config.Get(variable.DbDialect), config.Get(variable.DbUrl))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"event": "database",
		}).Error(err)
		return
	}
	db.LogMode(true)
	db.SetLogger(&gormrus.Logger{})

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
	userHandler := handler.NewUserHandler(userService, config)
	chatHandler := handler.NewChatHandler(chatService)

	// Create gin router
	router := handler.NewRouter(userHandler, chatHandler, config)

	// Start gin router
	if err := router.Run(":" + config.Get(variable.Port)); err != nil {
		logrus.WithFields(logrus.Fields{
			"event": "gin",
			"port":  config.Get(variable.Port),
		}).Error(err)
	}
}
