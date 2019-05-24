package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Credential struct {
	ID             uint      `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           string    `gorm:"unique" json:"name"`
	HashedPassword string    `json:"-"`
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
