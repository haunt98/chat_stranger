package models

import (
	"time"
)

type Credential struct {
	ID             int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Name           string `gorm:"unique"`
	HashedPassword string
}

type Authentication struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
