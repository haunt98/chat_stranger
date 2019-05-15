package handler

import (
	"net/http"
	"strconv"

	"github.com/1612180/chat_stranger/log"
	"github.com/1612180/chat_stranger/models"
	"github.com/1612180/chat_stranger/service"
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
	users, err := userHandler.service.FetchAll()
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func (userHandler *UserHandler) FindByID(c *gin.Context) {
	// Get id param
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
	} else {
		user, err := userHandler.service.FindByID(uint(id))
		if err != nil {
			log.ServerLog(err)
			c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
		} else {
			c.JSON(http.StatusOK, models.WrapSucceed{Succeed: true, Value: user})
		}
	}
}

func (userHandler *UserHandler) Create(c *gin.Context) {
	// Bind JSON
	var userPOST models.UserPOST
	err := c.ShouldBindJSON(&userPOST)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
	} else {
		err := userHandler.service.Create(&userPOST)
		if err != nil {
			log.ServerLog(err)
			c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
		} else {
			c.JSON(http.StatusOK, models.WrapSucceed{Succeed: true})
		}
	}
}

func (userHandler *UserHandler) UpdateByID(c *gin.Context) {
	// Get id param
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
	} else {
		// Bind JSON
		var userPOST models.UserPOST
		err := c.ShouldBindJSON(&userPOST)
		if err != nil {
			log.ServerLog(err)
			c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
		} else {
			err := userHandler.service.UpdateByID(uint(id), &userPOST)
			if err != nil {
				log.ServerLog(err)
				c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
			} else {
				c.JSON(http.StatusOK, models.WrapSucceed{Succeed: true})
			}
		}
	}
}

func (userHandler *UserHandler) DeleteByID(c *gin.Context) {
	// Get id param
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
	} else {
		err := userHandler.service.DeleteByID(uint(id))
		if err != nil {
			log.ServerLog(err)
			c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
		} else {
			c.JSON(http.StatusOK, models.WrapSucceed{Succeed: true})
		}
	}
}
