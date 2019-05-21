package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Credential struct {
	gorm.Model
	Name           string `gorm:"unique"`
	HashedPassword string `json:"-"`
}

type Authentication struct {
	Name     string
	Password string
}

type JWTClaims struct {
	ID   uint
	Role string
	jwt.StandardClaims
}
