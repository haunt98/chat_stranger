package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User belongs to Credential
type User struct {
	gorm.Model
	Credential   Credential
	CredentialID uint
}

type UserUpload struct {
	Authentication Authentication
}

func (userUpload *UserUpload) NewUser() (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(userUpload.Authentication.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := User{Credential: Credential{
		Name:           userUpload.Authentication.Name,
		HashedPassword: string(hashedPassword)},
	}

	return &user, nil
}
