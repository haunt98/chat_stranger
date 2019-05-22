package models

import (
	"github.com/jinzhu/gorm"
)

// Admin belongs to Credential
type Admin struct {
	gorm.Model   `json:"-"`
	Credential   Credential `json:"-"`
	CredentialID uint       `json:"-"`
	FullName     string     `json:"fullname"`
}

type AdminUpload struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
}
