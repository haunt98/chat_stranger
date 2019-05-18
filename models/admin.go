package models

import (
	"github.com/jinzhu/gorm"
)

// Admin belongs to Credential
type Admin struct {
	gorm.Model
	Credential   Credential
	CredentialID uint
}
