package handler

import (
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/jwt"
	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/1612180/chat_stranger/internal/service"
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	// bind json user
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "handler",
			"target": "user",
			"action": "sign up",
		}).Error(err)
		c.JSON(200, response.Create(102))
		return
	}

	if ok := h.userService.SignUp(&user); !ok {
		c.JSON(200, response.Create(101))
		return
	}

	// token
	s, ok := jwt.Create(jwt.SignClaims{
		ID:             user.ID,
		Role:           "user",
		StandardClaims: jwt2.StandardClaims{},
	}, viper.GetString(variable.JWTSecret))
	if !ok {
		c.JSON(200, response.Create(103))
		return
	}
	c.JSON(200, response.CreateWithData(100, s))
}

func (h *UserHandler) LogIn(c *gin.Context) {
	// bind json user
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "handler",
			"target": "user",
			"action": "log in",
		}).Error(err)
		c.JSON(200, response.Create(202))
		return
	}

	if ok := h.userService.LogIn(&user); !ok {
		c.JSON(200, response.Create(201))
		return
	}

	// token
	s, ok := jwt.Create(jwt.SignClaims{
		ID:             user.ID,
		Role:           "user",
		StandardClaims: jwt2.StandardClaims{},
	}, viper.GetString(variable.JWTSecret))
	if !ok {
		c.JSON(200, response.Create(203))
		return
	}
	c.JSON(200, response.CreateWithData(200, s))
}

func (h *UserHandler) Info(c *gin.Context) {
	// get user
	id, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	user, ok := h.userService.Info(id.(int))
	if !ok {
		c.JSON(200, response.Create(301))
		return
	}
	c.JSON(200, response.CreateWithData(300, user))
}

func (h *UserHandler) UpdateInfo(c *gin.Context) {
	// get user
	id, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	// bind json user
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "handler",
			"target": "user",
			"action": "update info",
		}).Error(err)
		c.JSON(200, response.Create(122))
		return
	}

	if ok := h.userService.UpdateInfo(id.(int), &user); !ok {
		c.JSON(200, response.Create(121))
		return
	}
	c.JSON(200, response.Create(120))
}

func (h *UserHandler) UpdatePassword(c *gin.Context) {
	// get user
	id, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	// bind json user
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "handler",
			"target": "user",
			"action": "update info",
		}).Error(err)
		c.JSON(200, response.Create(132))
		return
	}

	if ok := h.userService.UpdatePassword(id.(int), &user); !ok {
		c.JSON(200, response.Create(131))
		return
	}
	c.JSON(200, response.Create(130))
}
