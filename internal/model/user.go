package model

type User struct {
	ID           int    `json:"id"`
	FullName     string `json:"name"`
	Gender       string `json:"gender"`
	BirthYear    int    `json:"birth_year"`
	RegisterName string `json:"register_name" gorm:"-"`
	Password     string `json:"password" gorm:"-"`
	CredentialID int    `json:"-"`
}
