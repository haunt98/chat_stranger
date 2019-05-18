package main

import (
	"github.com/1612180/chat_stranger/config"
	"github.com/1612180/chat_stranger/handler"
	"github.com/1612180/chat_stranger/log"
	"github.com/1612180/chat_stranger/repository"
	"github.com/1612180/chat_stranger/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.ServerLog(err)
	}

	databaseConfig := config.NewDatabaseConfig()

	db, err := gorm.Open("postgres", databaseConfig.NewPostgresSource())
	if err != nil {
		log.ServerLog(err)
		return
	}
	defer db.Close()

	gin.DisableConsoleColor()
	router := gin.Default()

	credentialRepo := repository.NewCredentialRepoGorm(db)
	userRepo := repository.NewUserRepoGorm(db)
	userService := service.NewUserService(credentialRepo, userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.GET("/api/users", userHandler.FetchAll)
	router.GET("/api/users/:id", userHandler.FindByID)
	router.POST("/api/users", userHandler.Create)
	router.PUT("/api/users/:id/info", userHandler.UpdateInfoByID)
	router.PUT("/api/users/:id/password", userHandler.UpdatePasswordByID)
	router.DELETE("/api/users/:id", userHandler.DeleteByID)

	router.POST("/api/users/authenticate", userHandler.Authenticate)

	router.Run()
}
