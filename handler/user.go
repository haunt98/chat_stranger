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
		c.JSON(http.StatusInternalServerError, models.WrapSucceed{Succeed: false})
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func (userHandler *UserHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
	} else {
		user, err := userHandler.service.FindByID(uint(id))
		if err != nil {
			log.ServerLog(err)
			c.JSON(http.StatusInternalServerError, models.WrapSucceed{Succeed: false})
		} else {
			c.JSON(http.StatusOK, models.WrapSucceed{Succeed: true, Value: user})
		}
	}
}

func (userHandler *UserHandler) Create(c *gin.Context) {
	var userUpload models.UserUpload
	err := c.ShouldBindJSON(&userUpload)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
	} else {
		err := userHandler.service.Create(&userUpload)
		if err != nil {
			log.ServerLog(err)
			c.JSON(http.StatusInternalServerError, models.WrapSucceed{Succeed: false})
		} else {
			c.JSON(http.StatusOK, models.WrapSucceed{Succeed: true})
		}
	}
}

func (userHandler *UserHandler) UpdateByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
	} else {
		var userUpload models.UserUpload
		err := c.ShouldBindJSON(&userUpload)
		if err != nil {
			log.ServerLog(err)
			c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
		} else {
			err := userHandler.service.UpdateByID(uint(id), &userUpload)
			if err != nil {
				log.ServerLog(err)
				c.JSON(http.StatusInternalServerError, models.WrapSucceed{Succeed: false})
			} else {
				c.JSON(http.StatusOK, models.WrapSucceed{Succeed: true})
			}
		}
	}
}

func (userHandler *UserHandler) DeleteByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, models.WrapSucceed{Succeed: false})
	} else {
		err := userHandler.service.DeleteByID(uint(id))
		if err != nil {
			log.ServerLog(err)
			c.JSON(http.StatusInternalServerError, models.WrapSucceed{Succeed: false})
		} else {
			c.JSON(http.StatusOK, models.WrapSucceed{Succeed: true})
		}
	}
}

func (userHandler *UserHandler) Authenticate(c *gin.Context) {

}
