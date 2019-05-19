package models

import (
	"fmt"

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

func (authentication *Authentication) Authenticate(credential *Credential) error {
	if authentication.Name != credential.Name {
		return fmt.Errorf("Name or password is incorrect")
	}

	err := bcrypt.CompareHashAndPassword([]byte(credential.HashedPassword), []byte(authentication.Password))
	if err != nil {
		return fmt.Errorf("Name or password is incorrect")
	}

	return nil
}
