package models

import (
	"time"
)

// User belongs to Credential
type User struct {
	ID           int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Credential   Credential
	CredentialID int
	FullName     string
	Gender       string
	BirthYear    int
	Introduce    string
}

type UserUpload struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	FullName  string `json:"fullname"`
	Gender    string `json:"gender"`
	BirthYear int    `json:"birthyear"`
	Introduce string `json:"introduce"`
}

type UserDownload struct {
	ID        int    `json:"id"`
	FullName  string `json:"fullname"`
	Gender    string `json:"gender"`
	BirthYear int    `json:"birthyear"`
	Introduce string `json:"introduce"`
}
