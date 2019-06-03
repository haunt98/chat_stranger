package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Credential struct {
	ID             int       `json:"-"`
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
	ID   int
	Role string
	jwt.StandardClaims
}
