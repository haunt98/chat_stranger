package models

import (
	"time"

	"github.com/1612180/chat_stranger/internal/dtos"
	"golang.org/x/crypto/bcrypt"
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

func (user *User) FromRequest(userReq *dtos.UserRequest) (*User, []error) {
	var cre Credential
	cre.RegName = userReq.RegName

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return nil, errs
	}
	cre.HashedPassword = string(hashedPassword)

	user.Credential = cre
	user.FullName = userReq.FullName
	user.Gender = userReq.Gender
	user.BirthYear = userReq.BirthYear
	user.Introduce = userReq.Introduce

	return user, nil
}

func (user *User) UpdateFromRequest(userReq *dtos.UserRequest) *User {
	user.FullName = userReq.FullName
	user.Gender = userReq.Gender
	user.BirthYear = userReq.BirthYear
	user.Introduce = userReq.Introduce

	return user
}

func (user *User) ToResponse() *dtos.UserResponse {
	return &dtos.UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Gender:    user.Gender,
		BirthYear: user.BirthYear,
		Introduce: user.Introduce,
	}
}
