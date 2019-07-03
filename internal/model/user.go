package model

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Gender       string `json:"gender"`
	BirthYear    int    `json:"birth_year"`
	Password     string `json:"password" gorm:"-"`
	CredentialID int    `json:"-"`
}
