package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type SignClaims struct {
	ID   int
	Role string
	jwt.StandardClaims
}

func Create(claims jwt.Claims, secret string) (string, bool) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	s, err := token.SignedString([]byte(secret))
	if err != nil {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event": "jwt",
		}).Error("Failed to create token string")
		return "", false
	}
	return s, true
}

func Verify(s, secret string) (*SignClaims, bool) {
	token, err := jwt.ParseWithClaims(s, &SignClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if token == nil || !token.Valid {
		logrus.Error(err)
		logrus.WithFields(logrus.Fields{
			"event": "jwt",
		}).Error("Failed to verify token string")
		return nil, false
	}

	claims, ok := token.Claims.(*SignClaims)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"event": "jwt",
		}).Error("Failed to convert to claims")
		return nil, false
	}
	return claims, true
}
