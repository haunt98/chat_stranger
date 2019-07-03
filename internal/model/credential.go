package model

type Credential struct {
	ID             int
	Name           string `gorm:"unique"`
	HashedPassword string
}
