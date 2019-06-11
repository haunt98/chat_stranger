package main

import (
	"log"
	"net/http"

	"github.com/1612180/chat_stranger/internal/dtos"
	"github.com/1612180/chat_stranger/internal/handler"
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

	userService := service.NewUserService(credentialRepo, userRepo)
	adminService := service.NewAdminService(credentialRepo, adminRepo)
	roomService := service.NewRoomService(userRepo)

	adminService.Create(&dtos.AdminRequest{
		RegName:  viper.GetString("admin.regname"),
		Password: viper.GetString("admin.password"),
		FullName: viper.GetString("admin.fullname"),
	})

	userHandler := handler.NewUserHandler(userService)
	adminHandler := handler.NewAdminHandler(adminService)
	chatHandler := handler.NewChatHandler(roomService)

	gin.SetMode(viper.GetString("gin.mode"))
	gin.DisableConsoleColor()
	router := gin.Default()
	router.LoadHTMLGlob("./web/*.html")
	router.Static("/chat_stranger/web/script", "./web/script")

	web := router.Group("/chat_stranger/web")
	{
		web.GET("", func(c *gin.Context) {
			c.HTML(http.StatusOK, "home.html", gin.H{})
		})
		web.GET("/welcome", func(c *gin.Context) {
			c.HTML(http.StatusOK, "welcome.html", gin.H{})
		})
		web.GET("/chat", func(c *gin.Context) {
			c.HTML(http.StatusOK, "chat.html", gin.H{})
		})
	}

	r1 := router.Group("/chat_stranger/api")
	{
		auth := r1.Group("/auth")
		{
			auth.POST("/register", userHandler.Create)
			auth.POST("/login", userHandler.Authenticate)
			auth.POST("/login/admin", adminHandler.Authenticate)
		}
		r1.GET("/ws", chatHandler.WS)
	}

	r2 := router.Group("/chat_stranger/api/me", handler.VerifyRole("user"))
	{
		r2.GET("", userHandler.VerifyFind)
		r2.DELETE("", userHandler.VerifyDelete)
		r2.PUT("", userHandler.VerifyUpdateInfo)
		r2.GET("/room", chatHandler.FindRoom)
	}

	r3 := router.Group("/chat_stranger/api/users", handler.VerifyRole("admin"))
	{
		r3.GET("", userHandler.FetchAll)
		r3.GET("/:id", userHandler.Find)
		r3.POST("", userHandler.Create)
		r3.PUT("/:id", userHandler.UpdateInfo)
		r3.DELETE("/:id", userHandler.Delete)
	}

	r4 := router.Group("/chat_stranger/api/admins", handler.VerifyRole("admin"))
	{
		r4.GET("", adminHandler.FetchAll)
		r4.GET("/:id", adminHandler.Find)
		r4.POST("", adminHandler.Create)
		r4.PUT("/:id", adminHandler.UpdateInfo)
		r4.DELETE("/:id", adminHandler.Delete)
	}

	if err = router.Run(":" + viper.GetString("port")); err != nil {
		log.Println(err)
	}
}
