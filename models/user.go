package models

import (
	"time"
)

// User belongs to Credential
type User struct {
	ID           uint       `json:"-"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Credential   Credential `json:"-"`
	CredentialID uint       `json:"-"`
	Fullname     string     `json:"fullname"`
	Gender       string     `json:"gender"`
	Birthyear    int        `json:"birthyear"`
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
