package token

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type AccountClaims struct {
	ID   int
	Role string
	jwt.StandardClaims
}

func Create(claims jwt.Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	s, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func Verify(s, secret string) (AccountClaims, error) {
	token, err := jwt.ParseWithClaims(s, &AccountClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if token == nil || !token.Valid {
		return AccountClaims{}, err
	}

	accountClaims, ok := token.Claims.(*AccountClaims)
	if !ok {
		return AccountClaims{}, fmt.Errorf("parse to AccountClaims failed")
	}
	return *accountClaims, nil
}
