package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/1612180/chat_stranger/log"
	"github.com/1612180/chat_stranger/models"
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
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	res := Response(true, ":)")
	res["Users"] = users
	c.JSON(http.StatusOK, res)
}

func (userHandler *UserHandler) Find(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	user, errs := userHandler.service.Find(uint(id))
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	res := Response(true, ":)")
	res["User"] = user
	c.JSON(http.StatusOK, res)
}

func (userHandler *UserHandler) Create(c *gin.Context) {
	var userUpload models.UserUpload
	if err := c.ShouldBindJSON(&userUpload); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, "Username is already used"))
		return
	}

	id, errs := userHandler.service.Create(&userUpload)
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusBadRequest, Response(false, "Username is already used"))
		return
	}

	res := Response(true, "Register OK")
	res["UserID"] = id
	c.JSON(http.StatusOK, res)
}

func (userHandler *UserHandler) UpdateInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	var userUpload models.UserUpload
	if err = c.ShouldBindJSON(&userUpload); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	if errs := userHandler.service.UpdateInfo(uint(id), &userUpload); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, "Update Info OK"))
}

func (userHandler *UserHandler) UpdatePassword(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
	}

	var authentication models.Authentication
	if err = c.ShouldBindJSON(&authentication); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	if errs := userHandler.service.UpdatePassword(uint(id), &authentication); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, "Update Password OK"))
}

func (userHandler *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	if errs := userHandler.service.Delete(uint(id)); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, "Delete OK"))
}

func (userHandler *UserHandler) Authenticate(c *gin.Context) {
	var authentication models.Authentication
	if err := c.ShouldBindJSON(&authentication); err != nil {
		log.ServerLog(err)
		res := Response(false, "Username or password is incorrect")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user, errs := userHandler.service.Authenticate(&authentication)
	if len(errs) != 0 {
		log.ServerLogs(errs)
		res := Response(false, "Username or password is incorrect")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JWTClaims{
		user.ID,
		"User",
		jwt.StandardClaims{},
	})

	tokenStr, err := jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.ServerLog(err)
		res := Response(false, "Username or password is incorrect")
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := Response(true, "Login OK")
	res["Token"] = tokenStr
	c.JSON(http.StatusOK, res)
}

func (userHandler *UserHandler) VerifyDelete(c *gin.Context) {
	id, ok := c.Get("ID")
	if !ok {
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	if errs := userHandler.service.Delete(id.(uint)); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, "Delete OK"))
}

func (userHandler *UserHandler) VerifyUpdateInfo(c *gin.Context) {
	id, ok := c.Get("ID")
	if !ok {
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	var userUpload models.UserUpload
	if err := c.ShouldBindJSON(&userUpload); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	if errs := userHandler.service.UpdateInfo(id.(uint), &userUpload); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, "Update Info OK"))
}

func (userHandler *UserHandler) VerifyUpdatePassword(c *gin.Context) {
	id, ok := c.Get("ID")
	if !ok {
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	var authentication models.Authentication
	if err := c.ShouldBindJSON(&authentication); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	if errs := userHandler.service.UpdatePassword(id.(uint), &authentication); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, "Update Password OK"))
}
