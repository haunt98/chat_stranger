package models

import (
	"time"
)

// Admin belongs to Credential
type Admin struct {
	ID           int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Credential   Credential
	CredentialID int
	FullName     string
}

type AdminUpload struct {
	RegName  string `json:"regname"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
}

type AdminDownload struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
}
