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
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.ServerLog(err)
	}

	// Load database config
	databaseConfig := config.NewDatabaseConfig()

	// Open database connect
	db, err := gorm.Open("postgres", databaseConfig.NewPostgresSource())
	if err != nil {
		log.ServerLog(err)
		return
	}
	defer db.Close()

	// Disable gin color
	gin.DisableConsoleColor()
	
	router := gin.Default()

	userRepo := repository.NewGormUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.GET("/api/users", userHandler.FetchAll)
	router.GET("/api/users/:id", userHandler.FindByID)
	router.POST("/api/users", userHandler.Create)
	router.PUT("/api/users/:id", userHandler.UpdateByID)
	router.DELETE("/api/users/:id", userHandler.DeleteByID)

	router.Run()
}
