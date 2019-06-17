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
	Favorites    []*Favorite `gorm:"many2many:user_favorites;"`
	Rooms        []*Room     `gorm:"many2many:room_users;"`
}

func (user *User) FromRequest(req *dtos.UserRequest) (*User, []error) {
	var cre Credential
	cre.RegName = req.RegName

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		var errs []error
		errs = append(errs, err)
		return nil, errs
	}
	cre.HashedPassword = string(hashedPassword)

	user.Credential = cre
	user.FullName = req.FullName
	user.Gender = req.Gender
	user.BirthYear = req.BirthYear
	user.Introduce = req.Introduce

	return user, nil
}

func (user *User) UpdateFromRequest(req *dtos.UserRequest) *User {
	user.FullName = req.FullName
	user.Gender = req.Gender
	user.BirthYear = req.BirthYear
	user.Introduce = req.Introduce

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
