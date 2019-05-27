package main

import (
	"net/http"
	"os"

	"github.com/1612180/chat_stranger/config"
	"github.com/1612180/chat_stranger/handler"
	"github.com/1612180/chat_stranger/log"
	"github.com/1612180/chat_stranger/models"
	"github.com/1612180/chat_stranger/repository"
	"github.com/1612180/chat_stranger/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.ServerLog(err)
	}

	databaseConfig := config.NewDatabaseConfig()

	db, err := gorm.Open("postgres", databaseConfig.NewPostgresSource())
	if err != nil {
		log.ServerLog(err)
		return
	}
	defer db.Close()

	credentialRepo := repository.NewCredentialRepoGorm(db)
	userRepo := repository.NewUserRepoGorm(db)
	adminRepo := repository.NewAdminRepoGorm(db)

	adminRepo.Create(&models.AdminUpload{
		Name:     os.Getenv("ADMIN_NAME"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		FullName: os.Getenv("ADMIN_NAME"),
	})

	userService := service.NewUserService(credentialRepo, userRepo)
	adminService := service.NewAdminService(credentialRepo, adminRepo)

	userHandler := handler.NewUserHandler(userService)
	adminHandler := handler.NewAdminHandler(adminService)

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

		public.GET("/users/roomid", func(c *gin.Context) {
			m := make(map[string]interface{})
			m["roomid"] = hub.NewRoom()
			c.JSON(200, m)
		})
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

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	if err = router.Run(":" + PORT); err != nil {
		log.ServerLog(err)
	}
}
