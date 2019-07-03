package model

type Credential struct {
	ID             int    `json:"-"`
	RegisterName   string `json:"-" gorm:"unique"`
	HashedPassword string `json:"-"`
}
