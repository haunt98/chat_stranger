package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/1612180/chat_stranger/log"
	"github.com/1612180/chat_stranger/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyRole(Role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.ServerLog(fmt.Errorf(ResponseCode[407]))
			c.JSON(http.StatusForbidden, Response(407))
			c.Abort()
			return
		}

		splitAuthHeader := strings.Split(authHeader, "Bearer")
		tokenString := strings.TrimSpace(splitAuthHeader[1])

		if tokenString == "" {
			log.ServerLog(fmt.Errorf(ResponseCode[407]))
			c.JSON(http.StatusForbidden, Response(407))
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if token == nil || !token.Valid {
			log.ServerLog(err)
			c.JSON(http.StatusForbidden, Response(408))
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*models.JWTClaims); ok {
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
}
