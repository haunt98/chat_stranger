package handler

import (
	"github.com/1612180/chat_stranger/internal/pkg/env"
	"github.com/1612180/chat_stranger/pkg/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewRouter(userHandler *UserHandler, roomHandler *RoomHandler) *gin.Engine {
	// Load gin config
	gin.SetMode(viper.GetString(env.GinMode))

	router := gin.New()
	router.Use(ginrus.Logger(), gin.Recovery())
	router.LoadHTMLGlob("./web/*.html")
	router.Static("/chat_stranger/web/script", "./web/script")

	web := router.Group(env.WebPrefix)
	{
		web.GET("", func(c *gin.Context) {
			c.HTML(200, "home.html", gin.H{})
		})
		web.GET("/chat", func(c *gin.Context) {
			c.HTML(200, "chat.html", gin.H{})
		})
	}

	api := router.Group(env.APIPrefix)
	{
		api.POST("/auth/signup", userHandler.SignUp)
		api.POST("/auth/login", userHandler.LogIn)
	}

	roleUser := router.Group(env.APIPrefix, VerifyRole("user"))
	{
		roleUser.GET("/me", userHandler.Info)
		roleUser.GET("/chat/find", roomHandler.Find)
		roleUser.POST("/chat/join", roomHandler.Join)
		roleUser.POST("/chat/leave", roomHandler.Leave)
		roleUser.POST("/chat/send", roomHandler.SendMessage)
		roleUser.GET("/chat/receive", roomHandler.ReceiveMessage)
		roleUser.GET("/chat/is_free", roomHandler.IsUserFree)
		roleUser.GET("/chat/count_member", roomHandler.CountMember)
	}

	return router
}
