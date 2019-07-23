package model

import (
	"time"

	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Credential struct {
	ID             int
	RegisterName   string `gorm:"unique"`
	HashedPassword string
}

type User struct {
	ID        int    `json:"id,omitempty"`
	ShowName  string `json:"showname,omitempty"`
	Gender    string `json:"gender,omitempty"`
	BirthYear int    `json:"birthyear,omitempty"`

	CredentialID int `json:"-"`
}

type Room struct {
	ID int `json:"id"`
}

type Member struct {
	ID     int
	UserID int
	RoomID int
}

type Message struct {
	ID           int       `json:"-"`
	CreatedAt    time.Time `json:"createdat"`
	Body         string    `json:"body" gorm:"type:text"`
	RoomID       int       `json:"roomid"`
	UserID       int       `json:"userid"`
	UserShowName string    `json:"usershowname"`
}

func Migrate(db *gorm.DB) error {
	if viper.GetString(variable.DbMode) == variable.Debug {
		if err := db.DropTableIfExists(&Credential{}, &User{},
			&Room{}, &Member{}, &Message{}).Error; err != nil {
			return errors.Wrap(err, "migrate failed")
		}
	}

	if err := db.AutoMigrate(&Credential{}, &User{},
		&Room{}, &Member{}, &Message{}).Error; err != nil {
		return errors.Wrap(err, "migrate failed")
	}
	return nil
}
