package handler

import (
	"fmt"
	"github.com/1612180/chat_stranger/service"
	"net/http"
	"strings"

	"github.com/1612180/chat_stranger/log"
	"github.com/gin-gonic/gin"
)

func GetTokenHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf(ResponseCode[407])
	}

	splitAuthHeader := strings.Split(authHeader, "Bearer")
	tokenString := strings.TrimSpace(splitAuthHeader[1])

	if tokenString == "" {
		return "", fmt.Errorf(ResponseCode[407])
	}

	return tokenString, nil
}

func VerifyRole(Role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := GetTokenHeader(c)
		if err != nil {
			log.ServerLog(err)
			c.JSON(http.StatusForbidden, Response(407))
			c.Abort()
			return
		}

		claims, err := service.VerifyTokenString(tokenString)
		if err != nil {
			log.ServerLog(err)
			c.JSON(http.StatusForbidden, Response(408))
			c.Abort()
			return
		}

		if claims.Role != Role {
			log.ServerLog(fmt.Errorf(ResponseCode[409]))
			c.JSON(http.StatusForbidden, Response(409))
			c.Abort()
			return
		}

		c.Set("ID", claims.ID)
		c.Next()
	}
}
