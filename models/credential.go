package models

import (
	"github.com/jinzhu/gorm"
)

type Credential struct {
	gorm.Model
	Name           string
	HashedPassword string `json:"-"`
}

type Authentication struct {
	Name     string
	Password string
}
