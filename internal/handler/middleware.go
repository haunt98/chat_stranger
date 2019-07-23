package handler

import (
	"fmt"
	"strings"

	"github.com/1612180/chat_stranger/internal/pkg/config"
	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/1612180/chat_stranger/internal/pkg/token"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func getToken(c *gin.Context) (string, error) {
	header := c.GetHeader("Authorization")
	if header == "" {
		return "", fmt.Errorf("header auth empty")
	}

	headers := strings.Split(header, "Bearer")
	if len(headers) < 2 {
		return "", fmt.Errorf("bearer empty")
	}

	tkn := strings.TrimSpace(headers[1])
	if tkn == "" {
		return "", fmt.Errorf("token empty")
	}

	return tkn, nil
}

type Role struct {
	config config.Config
}

func (r *Role) Verify(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tkn, err := getToken(c)
		if err != nil {
			logrus.Error(errors.Wrap(err, "middleware verify role failed"))
			c.JSON(403, response.Create(999))
			c.Abort()
			return
		}

		accountClaims, err := token.Verify(tkn, r.config.Get(variable.JWTSecret))
		if err != nil {
			logrus.Error(errors.Wrap(err, "middleware verify role failed"))
			c.JSON(403, response.Create(999))
			c.Abort()
			return
		}

		if accountClaims.Role != role {
			logrus.Error(errors.Wrap(fmt.Errorf("role not allowed"), "middleware verify role failed"))
			c.JSON(403, response.Create(999))
			c.Abort()
			return
		}

		c.Set("userID", accountClaims.ID)
		c.Next()
	}
}
