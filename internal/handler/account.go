package handler

import (
	"fmt"

	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/1612180/chat_stranger/internal/pkg/valid"
	"github.com/1612180/chat_stranger/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type AccountHandler struct {
	accountService service.AccountService
}

func NewAccountHandler(accountService service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

type SignUpSubmit struct {
	RegisterName string `json:"registername,omitempty"`
	Password     string `json:"password,omitempty"`
	ShowName     string `json:"showname,omitempty"`
}

type LogInSubmit struct {
	RegisterName string `json:"registername,omitempty"`
	Password     string `json:"password,omitempty"`
}

type InfoSubmit struct {
	ShowName  string `json:"showname, omitempty"`
	Gender    string `json:"gender,omitempty"`
	BirthYear int    `json:"birthyear,omitempty"`
}

func (h *AccountHandler) SignUp(c *gin.Context) {
	var submit SignUpSubmit
	if err := c.ShouldBindJSON(&submit); err != nil {
		logrus.Error(errors.Wrap(err, "account handler: sign up failed"))
		c.JSON(400, response.Create(102))
		return
	}

	if err := valid.CheckSignUpSubmit(submit.ShowName, submit.RegisterName, submit.Password); err != nil {
		logrus.Error(errors.Wrap(err, "account handler: sign up failed"))
		c.JSON(400, response.CreateWithMessage(103, errors.Cause(err).Error()))
		return
	}

	tkn, err := h.accountService.SignUp(submit.ShowName, submit.RegisterName, submit.Password)
	if err != nil {
		logrus.Error(errors.Wrap(err, "account handler: sign up failed"))
		c.JSON(500, response.Create(101))
		return
	}
	c.JSON(200, response.CreateWithData(100, tkn))
}

func (h *AccountHandler) LogIn(c *gin.Context) {
	var submit LogInSubmit
	if err := c.ShouldBindJSON(&submit); err != nil {
		logrus.Error(errors.Wrap(err, "account handler: log in failed"))
		c.JSON(400, response.Create(202))
		return
	}

	if err := valid.CheckLogInSubmit(submit.RegisterName, submit.Password); err != nil {
		logrus.Error(errors.Wrap(err, "account handler: log in failed"))
		c.JSON(400, response.CreateWithMessage(203, errors.Cause(err).Error()))
		return
	}

	tkn, err := h.accountService.LogIn(submit.RegisterName, submit.Password)
	if err != nil {
		logrus.Error(errors.Wrap(err, "account handler: log in failed"))
		c.JSON(500, response.Create(201))
		return
	}
	c.JSON(200, response.CreateWithData(200, tkn))
}

func (h *AccountHandler) Info(c *gin.Context) {
	tempUserID, ok := c.Get("userID")
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("userID empty"), "account handler: info failed"))
		c.JSON(403, response.Create(999))
		return
	}

	userID, ok := tempUserID.(int)
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("type assertion failed"), "account handler: info failed"))
		c.JSON(403, response.Create(999))
		return
	}

	user, err := h.accountService.Info(userID)
	if err != nil {
		logrus.Error(errors.Wrap(err, "account handler: info failed"))
		c.JSON(500, response.Create(301))
		return
	}
	c.JSON(200, response.CreateWithData(300, user))
}

func (h *AccountHandler) UpdateInfo(c *gin.Context) {
	tempUserID, ok := c.Get("userID")
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("userID empty"), "account handler: update info failed"))
		c.JSON(403, response.Create(999))
		return
	}

	userID, ok := tempUserID.(int)
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("type assertion failed"), "account handler: update info failed"))
		c.JSON(403, response.Create(999))
		return
	}

	var submit InfoSubmit
	if err := c.ShouldBindJSON(&submit); err != nil {
		logrus.Error(errors.Wrap(err, "account handler: update info failed"))
		c.JSON(400, response.Create(122))
		return
	}

	if err := valid.CheckUpdateInfoSubmit(submit.ShowName, submit.Gender, submit.BirthYear); err != nil {
		logrus.Error(errors.Wrap(err, "account handler: update info failed"))
		c.JSON(400, response.CreateWithMessage(123, errors.Cause(err).Error()))
		return
	}

	user, err := h.accountService.UpdateInfo(userID, submit.ShowName, submit.Gender, submit.BirthYear)
	if err != nil {
		logrus.Error(errors.Wrap(err, "account handler: update info failed"))
		c.JSON(500, response.Create(121))
		return
	}
	c.JSON(200, response.CreateWithData(120, user))
}
