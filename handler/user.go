package handler

import (
	"fmt"
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
		res := Response(false, ":(")
		c.JSON(http.StatusInternalServerError, res)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func (userHandler *UserHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		res := Response(false, ":(")
		c.JSON(http.StatusBadRequest, res)
	} else {
		user, errs := userHandler.service.FindByID(uint(id))
		if len(errs) != 0 {
			log.ServerLogs(errs)
			res := Response(false, ":(")
			c.JSON(http.StatusInternalServerError, res)
		} else {
			res := Response(true, ":)")
			res["User"] = user
			c.JSON(http.StatusOK, res)
		}
	}
}

func (userHandler *UserHandler) Create(c *gin.Context) {
	var userUpload models.UserUpload
	err := c.ShouldBindJSON(&userUpload)
	if err != nil {
		log.ServerLog(err)
		res := Response(false, ":(")
		c.JSON(http.StatusBadRequest, res)
	} else {
		errs := userHandler.service.Create(&userUpload)
		if len(errs) != 0 {
			log.ServerLogs(errs)
			res := Response(false, ":(")
			c.JSON(http.StatusInternalServerError, res)
		} else {
			res := Response(true, ":)")
			c.JSON(http.StatusOK, res)
		}
	}
}

func (userHandler *UserHandler) UpdateInfoByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		res := Response(false, ":(")
		c.JSON(http.StatusBadRequest, res)
	} else {
		var userUpload models.UserUpload
		err := c.ShouldBindJSON(&userUpload)
		if err != nil {
			log.ServerLog(err)
			res := Response(false, ":(")
			c.JSON(http.StatusInternalServerError, res)
		} else {
			errs := userHandler.service.UpdateInfoByID(uint(id), &userUpload)
			if len(errs) != 0 {
				log.ServerLogs(errs)
				res := Response(false, ":(")
				c.JSON(http.StatusInternalServerError, res)
			} else {
				res := Response(true, ":)")
				c.JSON(http.StatusOK, res)
			}
		}
	}
}

func (userHandler *UserHandler) UpdatePasswordByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		res := Response(false, ":(")
		c.JSON(http.StatusBadRequest, res)
	} else {
		var userUpload models.UserUpload
		err := c.ShouldBindJSON(&userUpload)
		if err != nil {
			log.ServerLog(err)
			res := Response(false, ":(")
			c.JSON(http.StatusBadRequest, res)
		} else {
			errs := userHandler.service.UpdatePasswordByID(uint(id), &userUpload)
			if len(errs) != 0 {
				log.ServerLogs(errs)
				res := Response(false, ":(")
				c.JSON(http.StatusInternalServerError, res)
			} else {
				res := Response(true, ":)")
				c.JSON(http.StatusOK, res)
			}
		}
	}
}

func (userHandler *UserHandler) DeleteByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		res := Response(false, ":(")
		c.JSON(http.StatusBadRequest, res)
	} else {
		errs := userHandler.service.DeleteByID(uint(id))
		if len(errs) != 0 {
			log.ServerLogs(errs)
			res := Response(false, ":(")
			c.JSON(http.StatusInternalServerError, res)
		} else {
			res := Response(true, ":)")
			c.JSON(http.StatusOK, res)
		}
	}
}

func (userHandler *UserHandler) Authenticate(c *gin.Context) {
	var authentication models.Authentication
	err := c.ShouldBindJSON(&authentication)
	if err != nil {
		log.ServerLog(err)
		res := Response(false, ":(")
		c.JSON(http.StatusBadRequest, res)
	} else {
		errs := userHandler.service.Authenticate(&authentication)
		if len(errs) != 0 {
			log.ServerLogs(errs)
			res := Response(false, ":(")
			c.JSON(http.StatusBadRequest, res)
		} else {
			jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"Name": authentication.Name,
			})
			tokenStr, err := jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
			if err != nil {
				res := Response(false, ":(")
				c.JSON(http.StatusInternalServerError, res)
			} else {
				res := Response(true, "Login OK")
				res["Token"] = tokenStr
				c.JSON(http.StatusOK, res)
			}
		}
	}
}

func (userHandler *UserHandler) RequireAuthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Token")
		if tokenStr == "" {
			c.JSON(http.StatusBadRequest, Response(false, "Login require"))
			c.Abort()
		} else {
			token, error := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Error jwt")
				}
				return []byte(os.Getenv("SECRET_KEY")), nil
			})
			if error != nil {
				c.JSON(http.StatusBadRequest, Response(false, "Fuck"))
				c.Abort()
			}
			if token.Valid {
				c.Set("Decoded", token.Claims)
				c.Next()
			} else {
				c.JSON(http.StatusBadRequest, Response(false, "Fuck"))
				c.Abort()
			}
		}
	}
}
