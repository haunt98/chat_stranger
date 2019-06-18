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
	favRepo := repository.NewFavoriteRepoGorm(db)
	roomRepo := repository.NewRoomRepoGorm(db)

	userService := service.NewUserService(credentialRepo, userRepo)
	adminService := service.NewAdminService(credentialRepo, adminRepo)
	favService := service.NewFavoriteService(favRepo)
	roomService := service.NewRoomService(roomRepo, userRepo)

	adminService.Create(&dtos.AdminRequest{
		RegName:  viper.GetString("admin.regname"),
		Password: viper.GetString("admin.password"),
		FullName: viper.GetString("admin.fullname"),
	})

	userHandler := handler.NewUserHandler(userService)
	adminHandler := handler.NewAdminHandler(adminService)
	favHandler := handler.NewFavoriteHandler(favService)
	roomHandler := handler.NewRoomHandler(roomService)

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

	public := router.Group("/chat_stranger/api")
	{
		auth := public.Group("/auth")
		{
			auth.POST("/register", userHandler.Create)
			auth.POST("/login", userHandler.Authenticate)
			auth.POST("/login/admin", adminHandler.Authenticate)
		}
	}

	userRoute := router.Group("/chat_stranger/api", handler.VerifyRole("user"))
	{
		me := userRoute.Group("/me")
		{
			me.GET("", userHandler.VerifyFind)
			me.DELETE("", userHandler.VerifyDelete)
			me.PUT("", userHandler.VerifyUpdateInfo)
		}

		chat := userRoute.Group("/chat")
		{
			chat.GET("/empty", roomHandler.FindEmpty)
			chat.POST("/join", roomHandler.Join)
			chat.POST("/leave", roomHandler.Leave)
			chat.POST("/send", roomHandler.ReceiveMsg)
			chat.POST("/receive", roomHandler.SendMsg)
		}
	}

	adminRoute := router.Group("/chat_stranger/api", handler.VerifyRole("admin"))
	{
		RESTUser := adminRoute.Group("/users")
		{
			RESTUser.GET("", userHandler.FetchAll)
			RESTUser.GET("/:id", userHandler.Find)
			RESTUser.POST("", userHandler.Create)
			RESTUser.PUT("/:id", userHandler.UpdateInfo)
			RESTUser.DELETE("/:id", userHandler.Delete)
		}

		RESTAdmin := adminRoute.Group("/admins")
		{
			RESTAdmin.GET("", adminHandler.FetchAll)
			RESTAdmin.GET("/:id", adminHandler.Find)
			RESTAdmin.POST("", adminHandler.Create)
			RESTAdmin.PUT("/:id", adminHandler.UpdateInfo)
			RESTAdmin.DELETE("/:id", adminHandler.Delete)
		}

		RESTFav := adminRoute.Group("/favorites")
		{
			RESTFav.GET("", favHandler.FetchAll)
		}

		RESTRoom := adminRoute.Group("/rooms")
		{
			RESTRoom.GET("", roomHandler.FetchAll)
		}
	}

	if err = router.Run(":" + viper.GetString("port")); err != nil {
		log.Println(err)
	}
}
