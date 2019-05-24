package models

import (
	"time"
)

// Admin belongs to Credential
type Admin struct {
	ID           uint       `json:"-"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Credential   Credential `json:"-"`
	CredentialID uint       `json:"-"`
	Fullname     string     `json:"fullname"`
}

type AdminUpload struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
}
