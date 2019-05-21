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
	Gender       string
	BirthYear    int
	Introduce    string
}

type UserUpload struct {
	Authentication Authentication
	FullName       string
	Gender         string
	BirthYear      int
	Introduce      string
}
