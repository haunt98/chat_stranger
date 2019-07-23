package main

import (
	"github.com/1612180/chat_stranger/internal/handler"
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/config"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/1612180/chat_stranger/internal/pkg/viperwrap"
	"github.com/1612180/chat_stranger/internal/repository"
	"github.com/1612180/chat_stranger/internal/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load cfg
	viperwrap.Load(variable.ServiceName, variable.ConfigFile, variable.ConfigPath)
	cfg := config.NewConfig(variable.Viper)

	// Load database
	db, err := gorm.Open(cfg.Get(variable.DbDialect), cfg.Get(variable.DbUrl))
	if err != nil {
		logrus.Error(errors.Wrap(err, "database open failed"))
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			logrus.Error(errors.Wrap(err, "database close failed"))
			return
		}
	}()

	if err := model.Migrate(db); err != nil {
		logrus.Error(errors.Wrap(err, "database failed"))
	}

	// Load repository
	accountRepo := repository.NewAccountRepo(db)
	chatRepo := repository.NewChatRepo(db)

	// Load service
	accountService := service.NewAccountService(accountRepo, cfg)
	chatService := service.NewChatService(accountRepo, chatRepo)

	// Load handler
	accountHandler := handler.NewAccountHandler(accountService)
	chatHandler := handler.NewChatHandler(chatService)

	// CreateRoom gin router
	router := handler.NewRouter(accountHandler, chatHandler, cfg)

	// Start gin router
	if err := router.Run(":" + cfg.Get(variable.Port)); err != nil {
		logrus.Error(errors.Wrapf(err, "gin run port=%s failed", cfg.Get(variable.Port)))
	}
}
