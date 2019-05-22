package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Credential struct {
	gorm.Model     `json:"-"`
	Name           string `gorm:"unique" json:"name"`
	HashedPassword string `json:"-"`
}

type Authentication struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type JWTClaims struct {
	ID   uint
	Role string
	jwt.StandardClaims
}
