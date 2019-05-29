package service

import (
	"fmt"
	"github.com/1612180/chat_stranger/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func CreateTokenString(claims jwt.Claims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := jwtToken.SignedString([]byte(viper.GetString("jwt.secret_key")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyTokenString(tokenString string) (*models.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.secret_key")), nil
	})

	if token == nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.JWTClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("can not convert to JWT claims")
}
