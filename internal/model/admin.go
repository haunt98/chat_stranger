package model

type Admin struct {
	ID           int    `json:"id"`
	FullName     string `json:"full_name"`
	RegisterName string `json:"register_name" gorm:"-"`
	Password     string `json:"password" gorm:"-"`
	CredentialID int    `json:"-"`
}
