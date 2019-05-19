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
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	res := Response(true, ":)")
	res["Users"] = users
	c.JSON(http.StatusOK, res)
}

func (userHandler *UserHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	user, errs := userHandler.service.FindByID(uint(id))
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
	err := c.ShouldBindJSON(&userUpload)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	errs := userHandler.service.Create(&userUpload)
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, ":)"))
}

func (userHandler *UserHandler) UpdateInfoByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	var userUpload models.UserUpload
	err = c.ShouldBindJSON(&userUpload)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	errs := userHandler.service.UpdateInfoByID(uint(id), &userUpload)
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, ":)"))
}

func (userHandler *UserHandler) UpdatePasswordByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
	}

	var userUpload models.UserUpload
	err = c.ShouldBindJSON(&userUpload)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	errs := userHandler.service.UpdatePasswordByID(uint(id), &userUpload)
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, ":)"))
}

func (userHandler *UserHandler) DeleteByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.ServerLog(err)
		c.JSON(http.StatusBadRequest, Response(false, ":("))
		return
	}

	errs := userHandler.service.DeleteByID(uint(id))
	if len(errs) != 0 {
		log.ServerLogs(errs)
		c.JSON(http.StatusInternalServerError, Response(false, ":("))
		return
	}

	c.JSON(http.StatusOK, Response(true, ":)"))
}

func (userHandler *UserHandler) Authenticate(c *gin.Context) {
	var authentication models.Authentication
	err := c.ShouldBindJSON(&authentication)
	if err != nil {
		log.ServerLog(err)
		res := Response(false, ":(")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	errs := userHandler.service.Authenticate(&authentication)
	if len(errs) != 0 {
		log.ServerLogs(errs)
		res := Response(false, ":(")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Name": authentication.Name,
	})
	tokenStr, err := jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		res := Response(false, ":(")
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res := Response(true, "Login OK")
	res["Token"] = tokenStr
	c.JSON(http.StatusOK, res)
}

func (userHandler *UserHandler) RequireAuthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Token")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, Response(false, "Login require"))
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, Response(false, "Login require"))
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("Decoded", claims)
			c.Next()
		}
	}
}
