package handler

import (
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/configwrap"
	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/1612180/chat_stranger/internal/pkg/token"
	"github.com/1612180/chat_stranger/internal/pkg/valid"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/1612180/chat_stranger/internal/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	userService service.UserService
	config      configwrap.Config
}

func NewUserHandler(userService service.UserService, config configwrap.Config) *UserHandler {
	return &UserHandler{
		userService: userService,
		config:      config,
	}
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

	if ok, checkMsg := valid.CheckRegisterName(user.RegisterName); !ok {
		c.JSON(200, response.CreateWithMessage(141, checkMsg))
		return
	}

	if ok, checkMsg := valid.CheckPassword(user.Password); !ok {
		c.JSON(200, response.CreateWithMessage(141, checkMsg))
		return
	}

	if ok, checkMsg := valid.CheckFullName(user.FullName); !ok {
		c.JSON(200, response.CreateWithMessage(141, checkMsg))
		return
	}

	if ok := h.userService.SignUp(&user); !ok {
		c.JSON(200, response.Create(101))
		return
	}

	// token
	s, ok := token.Create(token.SignClaims{
		ID:             user.ID,
		Role:           "user",
		StandardClaims: jwt.StandardClaims{},
	}, h.config.Get(variable.JWTSecret))
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

	if ok, checkMsg := valid.CheckRegisterName(user.RegisterName); !ok {
		c.JSON(200, response.CreateWithMessage(141, checkMsg))
		return
	}

	if ok, checkMsg := valid.CheckPassword(user.Password); !ok {
		c.JSON(200, response.CreateWithMessage(141, checkMsg))
		return
	}

	if ok := h.userService.LogIn(&user); !ok {
		c.JSON(200, response.Create(201))
		return
	}

	// token
	s, ok := token.Create(token.SignClaims{
		ID:             user.ID,
		Role:           "user",
		StandardClaims: jwt.StandardClaims{},
	}, h.config.Get(variable.JWTSecret))
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
