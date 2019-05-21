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
		splitAuthHeader := strings.Split(authHeader, "Bearer")
		tokenString := strings.TrimSpace(splitAuthHeader[1])

		fmt.Println(authHeader)
		fmt.Println(splitAuthHeader)
		fmt.Println(tokenString)

		if tokenString == "" {
			log.ServerLog(fmt.Errorf("Token not found"))
			c.JSON(http.StatusForbidden, Response(false, "Login require"))
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &models.CredentialClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if !token.Valid {
			log.ServerLog(err)
			c.JSON(http.StatusForbidden, Response(false, "Login require"))
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*models.CredentialClaims); ok {
			if claims.Role != Role {
				log.ServerLog(fmt.Errorf("Role bad"))
				c.JSON(http.StatusForbidden, Response(false, "Login require"))
				c.Abort()
				return
			}

			c.Set("TokenName", claims.Name)
			c.Next()
		}
	}
}
