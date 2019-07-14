package handler

import (
	"strings"

	"github.com/1612180/chat_stranger/internal/pkg/configwrap"
	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/1612180/chat_stranger/internal/pkg/token"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

type Role struct {
	config configwrap.Config
}

func (r *Role) Verify(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		s, ok := getToken(c)
		if !ok {
			c.JSON(403, response.Create(999))
			c.Abort()
			return
		}

		signClaims, ok := token.Verify(s, r.config.Get(variable.JWTSecret))
		if !ok {
			c.JSON(403, response.Create(999))
			c.Abort()
			return
		}

		if signClaims.Role != role {
			c.JSON(403, response.Create(999))
			c.Abort()
			return
		}

		c.Set("userID", signClaims.ID)
		c.Next()
	}
}
