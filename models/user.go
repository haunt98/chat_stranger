package models

import (
	"github.com/jinzhu/gorm"
)

// User belongs to Credential
type User struct {
	gorm.Model   `json:"-"`
	Credential   Credential `json:"-"`
	CredentialID uint       `json:"-"`
	FullName     string     `json:"fullname"`
	Gender       string     `json:"gender"`
	BirthYear    int        `json:"birthyear"`
	Introduce    string     `json:"introduce"`
}

type UserUpload struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	FullName  string `json:"fullname"`
	Gender    string `json:"gender"`
	BirthYear int    `json:"birthyear"`
	Introduce string `json:"introduce"`
}
