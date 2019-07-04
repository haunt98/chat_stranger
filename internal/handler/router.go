package handler

import (
	"github.com/1612180/chat_stranger/internal/pkg/env"
	"github.com/1612180/chat_stranger/pkg/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewRouter(userHandler *UserHandler) *gin.Engine {
	// Load gin config
	gin.SetMode(viper.GetString(env.GinMode))

	router := gin.New()
	router.Use(ginrus.Logger(), gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello")
	})

	api := router.Group(env.APIPrefix)
	{
		api.POST("/auth/signup", userHandler.SignUp)
		api.POST("/auth/login", userHandler.LogIn)
	}

	return router
}
