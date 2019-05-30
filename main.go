package main

import (
	"github.com/1612180/chat_stranger/config"
	"github.com/1612180/chat_stranger/handler"
	"github.com/1612180/chat_stranger/log"
	"github.com/1612180/chat_stranger/models"
	"github.com/1612180/chat_stranger/repository"
	"github.com/1612180/chat_stranger/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	config.Init()

	databaseReader := config.NewMySQL()
	db, err := gorm.Open(databaseReader.GetDBMS(), databaseReader.GetSource())
	if err != nil {
		log.ServerLog(err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.ServerLog(err)
		}
	}()

	credentialRepo := repository.NewCredentialRepoGorm(db)
	userRepo := repository.NewUserRepoGorm(db)
	adminRepo := repository.NewAdminRepoGorm(db)
	adminRepo.Create(&models.AdminUpload{
		Name:     viper.GetString("default_admin.name"),
		Password: viper.GetString("default_admin.password"),
		FullName: viper.GetString("default_admin.fullname"),
	})

	userService := service.NewUserService(credentialRepo, userRepo)
	adminService := service.NewAdminService(credentialRepo, adminRepo)

	userHandler := handler.NewUserHandler(userService)
	adminHandler := handler.NewAdminHandler(adminService)

	gin.SetMode(viper.GetString("gin.mode"))
	gin.DisableConsoleColor()
	router := gin.Default()
	router.LoadHTMLGlob("./static/*.html")
	router.Static("/static/script", "./static/script")

	// Serve HTML
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})
	router.GET("/welcome_user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome_user.html", gin.H{})
	})
	router.GET("/chat/:roomid", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat.html", gin.H{})
	})

	hub := handler.NewHub()
	router.GET("ws", hub.ChatHandler)

	public := router.Group("/api/public")
	{
		public.POST("/users/register", userHandler.Create)
		public.POST("/users/authenticate", userHandler.Authenticate)
		public.POST("/admins/authenticate", adminHandler.Authenticate)

		public.GET("/users/roomid", hub.GetRoom)
	}

	roleUser := router.Group("/api/me", handler.VerifyRole("User"))
	{
		roleUser.GET("", userHandler.VerifyFind)
		roleUser.DELETE("", userHandler.VerifyDelete)
		roleUser.PUT("/info", userHandler.VerifyUpdateInfo)
		roleUser.PUT("/password", userHandler.VerifyUpdatePassword)
	}

	roleAdmin := router.Group("/api/me", handler.VerifyRole("Admin"))
	{
		roleAdmin.GET("/users", userHandler.FetchAll)
		roleAdmin.GET("/users/:id", userHandler.Find)
		roleAdmin.POST("/users", userHandler.Create)
		roleAdmin.PUT("/users/:id/info", userHandler.UpdateInfo)
		roleAdmin.PUT("/users/:id/password", userHandler.UpdatePassword)
		roleAdmin.DELETE("/users/:id", userHandler.Delete)
		roleAdmin.GET("/admins", adminHandler.FetchAll)
		roleAdmin.GET("/admins/:id", adminHandler.Find)
		roleAdmin.POST("/admins", adminHandler.Create)
		roleAdmin.PUT("/admins/:id/info", adminHandler.UpdateInfo)
		roleAdmin.PUT("/admins/:id/password", adminHandler.UpdatePassword)
		roleAdmin.DELETE("/admins/:id", adminHandler.Delete)
	}

	if err = router.Run(":" + viper.GetString("port")); err != nil {
		log.ServerLog(err)
	}
}
