package handler

import (
	"net/http"
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
		c.JSON(http.StatusOK, Response(402))
		return
	}

	res := Response(200)
	res["admins"] = admins
	c.JSON(http.StatusOK, res)
}

func (adminHandler *AdminHandler) Find(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(401))
		return
	}

	admin, errs := adminHandler.service.Find(uint(id))
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(403))
		return
	}

	res := Response(201)
	res["admin"] = admin
	c.JSON(http.StatusOK, res)
}

func (adminHandler *AdminHandler) Create(c *gin.Context) {
	var adminUpload models.AdminUpload
	if err := c.ShouldBindJSON(&adminUpload); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	id, errs := adminHandler.service.Create(&adminUpload)
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(401))
		return
	}

	res := Response(205)
	res["adminid"] = id
	c.JSON(http.StatusOK, res)
}

func (adminHandler *AdminHandler) UpdateInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(401))
		return
	}

	var adminUpload models.AdminUpload
	if err = c.ShouldBindJSON(&adminUpload); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	if errs := adminHandler.service.UpdateInfo(uint(id), &adminUpload); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(202))
}

func (adminHandler *AdminHandler) UpdatePassword(c *gin.Context) {
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

	if errs := adminHandler.service.UpdatePassword(uint(id), &authentication); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(203))
}

func (adminHandler *AdminHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(401))
		return
	}

	if errs := adminHandler.service.Delete(uint(id)); len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(204))
}

func (adminHandler *AdminHandler) Authenticate(c *gin.Context) {
	var authentication models.Authentication
	if err := c.ShouldBindJSON(&authentication); err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	admin, errs := adminHandler.service.Authenticate(&authentication)
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusOK, Response(405))
		return
	}

	tokenString, err := service.CreateTokenString(models.JWTClaims{
		ID:             admin.ID,
		Role:           "Admin",
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
