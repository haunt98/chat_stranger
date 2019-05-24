package main

import (
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

	router.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	router.GET("/welcome_user", func(c *gin.Context) {
		c.HTML(200, "welcome_user.html", gin.H{})
	})

	public := router.Group("/api/public")
	{
		public.POST("/users/register", userHandler.Create)
		public.POST("/users/authenticate", userHandler.Authenticate)
		public.POST("/admins/authenticate", adminHandler.Authenticate)
	}

	privateForUser := router.Group("/api/privateForUser")
	privateForUser.Use(handler.VerifyRole("User"))
	{
		privateForUser.GET("", userHandler.VerifyFind)
		privateForUser.DELETE("", userHandler.VerifyDelete)
		privateForUser.PUT("/info", userHandler.VerifyUpdateInfo)
		privateForUser.PUT("/password", userHandler.VerifyUpdatePassword)
	}

	privateForAdmin := router.Group("/api/privateForAdmin")
	privateForAdmin.Use(handler.VerifyRole("Admin"))
	{
		privateForAdmin.GET("/users", userHandler.FetchAll)
		privateForAdmin.GET("/users/:id", userHandler.Find)
		privateForAdmin.POST("/users", userHandler.Create)
		privateForAdmin.PUT("/users/:id/info", userHandler.UpdateInfo)
		privateForAdmin.PUT("/users/:id/password", userHandler.UpdatePassword)
		privateForAdmin.DELETE("/users/:id", userHandler.Delete)
		privateForAdmin.GET("/admins", adminHandler.FetchAll)
		privateForAdmin.GET("/admins/:id", adminHandler.Find)
		privateForAdmin.POST("/admins", adminHandler.Create)
		privateForAdmin.PUT("/admins/:id/info", adminHandler.UpdateInfo)
		privateForAdmin.PUT("/admins/:id/password", adminHandler.UpdatePassword)
		privateForAdmin.DELETE("/admins/:id", adminHandler.Delete)
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	_ = router.Run(":" + PORT)
}
