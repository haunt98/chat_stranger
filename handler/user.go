package handler

import (
	"net/http"
	"strconv"

	"github.com/1612180/chat_stranger/models"
	"github.com/1612180/chat_stranger/pkg/log"
	"github.com/1612180/chat_stranger/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (userHandler *UserHandler) FetchAll(c *gin.Context) {
	users, errs := userHandler.service.FetchAll()
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(402))
		return
	}

	res := Response(200)
	res["users"] = users
	c.JSON(http.StatusOK, res)
}

func (userHandler *UserHandler) Find(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(401))
		return
	}

	user, errs := userHandler.service.Find(uint(id))
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(403))
		return
	}

	res := Response(201)
	res["user"] = user
	c.JSON(http.StatusOK, res)
}

func (userHandler *UserHandler) Create(c *gin.Context) {
	var userUpload models.UserUpload
	if err := c.ShouldBindJSON(&userUpload); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	id, errs := userHandler.service.Create(&userUpload)
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(404))
		return
	}

	res := Response(205)
	res["userid"] = id
	c.JSON(http.StatusOK, res)
}

func (userHandler *UserHandler) UpdateInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(401))
		return
	}

	var userUpload models.UserUpload
	if err = c.ShouldBindJSON(&userUpload); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	if errs := userHandler.service.UpdateInfo(uint(id), &userUpload); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(202))
}

func (userHandler *UserHandler) UpdatePassword(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(401))
	}

	var authentication models.Authentication
	if err = c.ShouldBindJSON(&authentication); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	if errs := userHandler.service.UpdatePassword(uint(id), &authentication); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(203))
}

func (userHandler *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(401))
		return
	}

	if errs := userHandler.service.Delete(uint(id)); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(204))
}

func (userHandler *UserHandler) Authenticate(c *gin.Context) {
	var authentication models.Authentication
	if err := c.ShouldBindJSON(&authentication); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	user, errs := userHandler.service.Authenticate(&authentication)
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(405))
		return
	}

	tokenString, err := service.CreateTokenString(models.JWTClaims{
		ID:             user.ID,
		Role:           "User",
		StandardClaims: jwt.StandardClaims{},
	})
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusInternalServerError, Response(500))
		return
	}

	res := Response(206)
	res["token"] = tokenString
	c.JSON(http.StatusOK, res)
}

func (userHandler *UserHandler) VerifyFind(c *gin.Context) {
	id, ok := c.Get("ID")
	if !ok {
		c.JSON(http.StatusBadRequest, Response(501))
		return
	}

	user, errs := userHandler.service.Find(id.(uint))
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(403))
		return
	}

	res := Response(201)
	res["user"] = user
	c.JSON(http.StatusOK, res)
}

func (userHandler *UserHandler) VerifyDelete(c *gin.Context) {
	id, ok := c.Get("ID")
	if !ok {
		c.JSON(http.StatusBadRequest, Response(501))
		return
	}

	if errs := userHandler.service.Delete(id.(uint)); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(204))
}

func (userHandler *UserHandler) VerifyUpdateInfo(c *gin.Context) {
	id, ok := c.Get("ID")
	if !ok {
		c.JSON(http.StatusInternalServerError, Response(501))
		return
	}

	var userUpload models.UserUpload
	if err := c.ShouldBindJSON(&userUpload); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	if errs := userHandler.service.UpdateInfo(id.(uint), &userUpload); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(202))
}

func (userHandler *UserHandler) VerifyUpdatePassword(c *gin.Context) {
	id, ok := c.Get("ID")
	if !ok {
		c.JSON(http.StatusInternalServerError, Response(501))
		return
	}

	var authentication models.Authentication
	if err := c.ShouldBindJSON(&authentication); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	if errs := userHandler.service.UpdatePassword(id.(uint), &authentication); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(203))
}
