package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Credential struct {
	gorm.Model
	Name           string
	HashedPassword string `json:"-"`
}

type Authentication struct {
	Name     string
	Password string
}

func (authentication *Authentication) NewCredential() (*Credential, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(authentication.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	credential := Credential{
		Name:           authentication.Name,
		HashedPassword: string(hashedPassword),
	}

	return &credential, nil
}
