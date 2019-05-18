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
	users, errs := userHandler.service.FetchAll()
	if len(errs) != 0 {
		log.ServerLogs(errs)
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
		user, errs := userHandler.service.FindByID(uint(id))
		if len(errs) != 0 {
			log.ServerLogs(errs)
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
		status, errs := userHandler.service.Create(&userUpload)
		if status == false {
			log.ServerLogs(errs)
			c.JSON(http.StatusInternalServerError, models.WrapSucceed{Succeed: false})
		} else {
			c.JSON(http.StatusOK, models.WrapSucceed{Succeed: true})
		}
	}
}

func (userHandler *UserHandler) UpdateInfoByID(c *gin.Context) {
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
			status, errs := userHandler.service.UpdateInfoByID(uint(id), &userUpload)
			if status == false {
				log.ServerLogs(errs)
				c.JSON(http.StatusInternalServerError, models.WrapSucceed{Succeed: false})
			} else {
				c.JSON(http.StatusOK, models.WrapSucceed{Succeed: true})
			}
		}
	}
}

func (userHandler *UserHandler) UpdatePasswordByID(c *gin.Context) {
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
			status, errs := userHandler.service.UpdatePasswordByID(uint(id), &userUpload)
			if status == false {
				log.ServerLogs(errs)
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
		status, errs := userHandler.service.DeleteByID(uint(id))
		if status == false {
			log.ServerLogs(errs)
			c.JSON(http.StatusInternalServerError, models.WrapSucceed{Succeed: false})
		} else {
			c.JSON(http.StatusOK, models.WrapSucceed{Succeed: true})
		}
	}
}

func (userHandler *UserHandler) Authenticate(c *gin.Context) {

}
