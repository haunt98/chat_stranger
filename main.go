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
	adminRepo := repository.NewAdminRepoGorm(db)

	userService := service.NewUserService(credentialRepo, userRepo)
	adminService := service.NewAdminService(credentialRepo, adminRepo)

	userHandler := handler.NewUserHandler(userService)
	adminHandler := handler.NewAdminHandler(adminService)

	RESTUser := router.Group("/api/users")
	//RESTUser.Use(handler.VerifyRole("Admin"))
	{
		RESTUser.GET("", userHandler.FetchAll)
		RESTUser.GET("/:id", userHandler.Find)
		RESTUser.POST("", userHandler.Create)
		RESTUser.PUT("/:id/info", userHandler.UpdateInfo)
		RESTUser.PUT("/:id/password", userHandler.UpdatePassword)
		RESTUser.DELETE("/:id", userHandler.Delete)
	}

	RESTAdmin := router.Group("/api/admins")
	{
		RESTAdmin.GET("", adminHandler.FetchAll)
		RESTAdmin.GET("/:id", adminHandler.Find)
		RESTAdmin.POST("", adminHandler.Create)
		RESTAdmin.PUT("/:id/info", adminHandler.UpdateInfo)
		RESTAdmin.PUT("/:id/password", adminHandler.UpdatePassword)
		RESTAdmin.DELETE("/:id", adminHandler.Delete)
	}

	router.POST("/api/users/authenticate", userHandler.Authenticate)

	router.Run()
}
