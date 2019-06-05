package main

import (
	"log"
	"net/http"

	"github.com/1612180/chat_stranger/internal/handler"
	"github.com/1612180/chat_stranger/internal/models"
	"github.com/1612180/chat_stranger/internal/repository"
	"github.com/1612180/chat_stranger/internal/service"
	"github.com/1612180/chat_stranger/pkg/configutils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func main() {
	log.SetPrefix("[Server] ")
	configutils.LoadConfiguration("chat_stranger", "config", "configs")

	db, err := gorm.Open(viper.GetString("db.dialect"), viper.GetString("db.url"))
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
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
	roomService := service.NewRoomService(userRepo)

	userHandler := handler.NewUserHandler(userService)
	adminHandler := handler.NewAdminHandler(adminService)
	chatHandler := handler.NewChatHandler(roomService)

	gin.SetMode(viper.GetString("gin.mode"))
	gin.DisableConsoleColor()
	router := gin.Default()
	router.LoadHTMLGlob("./web/*.html")
	router.Static("/web/script", "./web/script")

	// Serve HTML
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})
	router.GET("/welcome_user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome_user.html", gin.H{})
	})
	router.GET("/chat/:rid", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat.html", gin.H{})
	})

	public := router.Group("/api/public")
	{
		public.POST("/users/register", userHandler.Create)
		public.POST("/users/authenticate", userHandler.Authenticate)
		public.POST("/admins/authenticate", adminHandler.Authenticate)
		public.GET("/ws", chatHandler.WS)
	}

	roleUser := router.Group("/api/me", handler.VerifyRole("User"))
	{
		roleUser.GET("", userHandler.VerifyFind)
		roleUser.DELETE("", userHandler.VerifyDelete)
		roleUser.PUT("/info", userHandler.VerifyUpdateInfo)
		roleUser.PUT("/password", userHandler.VerifyUpdatePassword)
		roleUser.GET("/room", chatHandler.FindRoom)
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
		log.Println(err)
	}
}
