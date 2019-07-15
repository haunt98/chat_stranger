package token

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
	if err := claims.Valid(); err != nil {
		logrus.WithFields(logrus.Fields{
			"module": "jwt",
		}).Error(err)
		return "", false
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	s, err := token.SignedString([]byte(secret))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"module": "jwt",
		}).Error(err)
		return "", false
	}
	return s, true
}

func Verify(s, secret string) (*SignClaims, bool) {
	token, err := jwt.ParseWithClaims(s, &SignClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if token == nil || !token.Valid {
		logrus.WithFields(logrus.Fields{
			"module": "jwt",
		}).Error(err)
		return nil, false
	}

	claims, ok := token.Claims.(*SignClaims)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"module": "jwt",
		}).Error("Failed to convert to claims")
		return nil, false
	}
	return claims, true
}
