package models

import (
	"github.com/jinzhu/gorm"
)

// User belongs to Credential
type User struct {
	gorm.Model
	Credential   Credential
	CredentialID uint
	FullName     string
}

type UserUpload struct {
	Authentication Authentication
	FullName       string
}
