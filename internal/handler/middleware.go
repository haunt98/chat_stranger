package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/1612180/chat_stranger/internal/service"
	"github.com/gin-gonic/gin"
)

func GetTokenHeader(c *gin.Context) (string, error) {
	header := c.GetHeader("Authorization")
	if header == "" {
		return "", fmt.Errorf(ResponseCode[407])
	}

	headers := strings.Split(header, "Bearer")
	if len(headers) < 2 {
		return "", fmt.Errorf(ResponseCode[407])
	}

	s := strings.TrimSpace(headers[1])

	if s == "" {
		return "", fmt.Errorf(ResponseCode[407])
	}

	return s, nil
}

func VerifyRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		s, err := GetTokenHeader(c)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusForbidden, Response(407))
			c.Abort()
			return
		}

		claims, err := service.VerifyTokenString(s)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusForbidden, Response(408))
			c.Abort()
			return
		}

		if claims.Role != role {
			log.Println(fmt.Errorf(ResponseCode[409]))
			c.JSON(http.StatusForbidden, Response(409))
			c.Abort()
			return
		}

		c.Set("ID", claims.ID)
		c.Next()
	}
}