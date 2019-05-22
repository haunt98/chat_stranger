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

type AdminHandler struct {
	service *service.AdminService
}

func NewAdminHandler(service *service.AdminService) *AdminHandler {
	return &AdminHandler{
		service: service,
	}
}

func (adminHandler *AdminHandler) FetchAll(c *gin.Context) {
	admins, errs := adminHandler.service.FetchAll()
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	res := Response(true, ":)")
	res["admins"] = admins
	c.JSON(http.StatusOK, res)
}

func (adminHandler *AdminHandler) Find(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	admin, errs := adminHandler.service.Find(uint(id))
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	res := Response(true, ":)")
	res["admin"] = admin
	c.JSON(http.StatusOK, res)
}

func (adminHandler *AdminHandler) Create(c *gin.Context) {
	var adminUpload models.AdminUpload
	if err := c.ShouldBindJSON(&adminUpload); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, "Username is already used"))
		return
	}

	id, errs := adminHandler.service.Create(&adminUpload)
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusBadRequest, Response(false, "Username is already used"))
		return
	}

	res := Response(true, "Register OK")
	res["adminid"] = id
	c.JSON(http.StatusOK, res)
}

func (adminHandler *AdminHandler) UpdateInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	var adminUpload models.AdminUpload
	if err = c.ShouldBindJSON(&adminUpload); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	if errs := adminHandler.service.UpdateInfo(uint(id), &adminUpload); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, "Update Info OK"))
}

func (adminHandler *AdminHandler) UpdatePassword(c *gin.Context) {
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

	if errs := adminHandler.service.UpdatePassword(uint(id), &authentication); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, "Update Password OK"))
}

func (adminHandler *AdminHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	if errs := adminHandler.service.Delete(uint(id)); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, "Delete OK"))
}

func (adminHandler *AdminHandler) Authenticate(c *gin.Context) {
	var authentication models.Authentication
	if err := c.ShouldBindJSON(&authentication); err != nil {
		log.ServerLog(err)
		res := Response(false, "Username or password is incorrect")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	admin, errs := adminHandler.service.Authenticate(&authentication)
	if len(errs) != 0 {
		log.ServerLogs(errs)
		res := Response(false, "Username or password is incorrect")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JWTClaims{
		admin.ID,
		"Admin",
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
	res["token"] = tokenStr
	c.JSON(http.StatusOK, res)
}
