package models

import (
	"github.com/jinzhu/gorm"
)

// Admin save admin information
type Admin struct {
	gorm.Model
	Credential Credential
}
