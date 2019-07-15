package handler

import (
	"github.com/1612180/chat_stranger/internal/pkg/configwrap"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/1612180/chat_stranger/pkg/ginrus"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewRouter(userHandler *UserHandler, chatHandler *ChatHandler, config configwrap.Config) *gin.Engine {
	// Load gin config
	gin.SetMode(viper.GetString(variable.GinMode))

	router := gin.New()
	router.Use(ginrus.Logger(), gin.Recovery())

	// add html when not test
	if config.Get(variable.ConfigMode) != variable.TestMode {
		router.LoadHTMLGlob(variable.HTMLGlob)
		router.Static(variable.StaticRelative, variable.StaticRoot)

		web := router.Group(variable.WebPrefix)
		{
			web.GET("", func(c *gin.Context) {
				c.HTML(200, "home.html", gin.H{})
			})
			web.GET("/chat", func(c *gin.Context) {
				c.HTML(200, "chat.html", gin.H{})
			})
		}
	}

	api := router.Group(variable.APIPrefix)
	{
		api.POST("/auth/signup", userHandler.SignUp)
		api.POST("/auth/login", userHandler.LogIn)
	}

	role := Role{config: config}
	roleUser := router.Group(variable.APIPrefix, role.Verify(variable.UserRole))
	{
		roleUser.GET("/me", userHandler.Info)
		roleUser.PUT("/me", userHandler.UpdateInfo)

		roleUser.GET("/chat/find", chatHandler.Find)
		roleUser.POST("/chat/join", chatHandler.Join)
		roleUser.POST("/chat/leave", chatHandler.Leave)
		roleUser.POST("/chat/send", chatHandler.SendMessage)
		roleUser.GET("/chat/receive", chatHandler.ReceiveMessage)
		roleUser.GET("/chat/is_free", chatHandler.IsUserFree)
		roleUser.GET("/chat/count_member", chatHandler.CountMember)
	}
	return router
}
