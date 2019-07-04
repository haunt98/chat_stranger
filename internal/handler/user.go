package handler

import (
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/1612180/chat_stranger/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event":  "handler",
			"target": "user",
		}).Error("Failed to bind json when sign up")
		c.JSON(200, response.Create(999))
		return
	}

	if ok := h.userService.SignUp(&user); !ok {
		c.JSON(200, response.Create(999))
		return
	}
	c.JSON(200, response.Create(1))
}

func (h *UserHandler) LogIn(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event":  "handler",
			"target": "user",
		}).Error("Failed to bind json when log in")
		c.JSON(200, response.Create(999))
		return
	}

	if ok := h.userService.LogIn(&user); !ok {
		c.JSON(200, response.Create(999))
		return
	}
	c.JSON(200, response.Create(1))
}
