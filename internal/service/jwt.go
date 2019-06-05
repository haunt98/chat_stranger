package service

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type JWTClaims struct {
	ID   int
	Role string
	jwt.StandardClaims
}

func CreateTokenString(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	s, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return "", err
	}

	return s, nil
}

func VerifyTokenString(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.secret")), nil
	})

	if token == nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("can not convert to JWT claims")
}
