package handler

import (
	"github.com/1612180/chat_stranger/internal/pkg/config"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/gin-gonic/gin"
)

func NewRouter(accountHandler *AccountHandler, chatHandler *ChatHandler, cfg config.Config) *gin.Engine {
	gin.SetMode(cfg.Get(variable.GinMode))
	router := gin.New()

	// production
	if cfg.Get(variable.Config) != variable.Test {
		router.Use(gin.Logger(), gin.Recovery())

		router.LoadHTMLGlob(variable.HTMLGlob)
		router.Static(variable.StaticRelativeScript, variable.StaticScript)
		router.Static(variable.StaticRelativeStyle, variable.StaticStyle)
		router.Static(variable.StaticRelativeImg, variable.StaticImg)

		web := router.Group(variable.WebPrefix)
		{
			web.GET("", func(c *gin.Context) {
				c.HTML(200, "home.html", gin.H{})
			})
			web.GET("/chat", func(c *gin.Context) {
				c.HTML(200, "chat.html", gin.H{})
			})
		}

		// redirect
		router.NoRoute(func(c *gin.Context) {
			c.Redirect(301, variable.WebPrefix)
		})
	}

	api := router.Group(variable.APIPrefix)
	{
		api.POST("/auth/signup", accountHandler.SignUp)
		api.POST("/auth/login", accountHandler.LogIn)
	}

	role := Role{config: cfg}
	roleUser := router.Group(variable.APIPrefix, role.Verify(variable.UserRole))
	{
		roleUser.GET("/me", accountHandler.Info)
		roleUser.PUT("/me", accountHandler.UpdateInfo)

		roleUser.GET("/chat/find", chatHandler.FindRoom)
		roleUser.POST("/chat/join", chatHandler.Join)
		roleUser.POST("/chat/leave", chatHandler.Leave)
		roleUser.POST("/chat/send", chatHandler.SendMessage)
		roleUser.GET("/chat/receive", chatHandler.ReceiveMessage)
		roleUser.GET("/chat/is_free", chatHandler.IsUserFree)
		roleUser.GET("/chat/count_member", chatHandler.CountMember)
	}
	return router
}
