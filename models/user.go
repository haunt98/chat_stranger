package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username     string
	PasswordHash string
}

// Struct for POST /users
type UserPOST struct {
	Username string
	Password string
}

func (up *UserPOST) NewUser() (*User, error) {
	// Use bcrypt to hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(up.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	} else {
		m := User{Username: up.Username, PasswordHash: string(passwordHash)}
		return &m, nil
	}
}
