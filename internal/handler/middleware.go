package handler

import (
	"strings"

	"github.com/1612180/chat_stranger/internal/pkg/env"
	"github.com/1612180/chat_stranger/internal/pkg/jwt"
	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func getToken(c *gin.Context) (string, bool) {
	header := c.GetHeader("Authorization")
	if header == "" {
		logrus.WithFields(logrus.Fields{
			"target": "header",
		}).Error("Header not found")
		return "", false
	}

	headers := strings.Split(header, "Bearer")
	if len(headers) < 2 {
		logrus.WithFields(logrus.Fields{
			"target": "header",
		}).Error("Bearer not found")
		return "", false
	}

	s := strings.TrimSpace(headers[1])
	if s == "" {
		logrus.WithFields(logrus.Fields{
			"target": "header",
		}).Error("Token not found")
		return "", false
	}

	return s, true
}

func VerifyRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		s, ok := getToken(c)
		if !ok {
			c.JSON(403, response.Create(9))
			c.Abort()
			return
		}

		signClaims, ok := jwt.Verify(s, viper.GetString(env.JWTSecret))
		if !ok {
			c.JSON(403, response.Create(9))
			c.Abort()
			return
		}

		if signClaims.Role != role {
			c.JSON(403, response.Create(9))
			c.Abort()
			return
		}

		c.Set("userID", signClaims.ID)
		c.Next()
	}
}
