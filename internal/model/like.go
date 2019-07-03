package model

type Like struct {
	UserID  int `json:"user_id" gorm:"primary_key"`
	HobbyID int `json:"hobby_id" gorm:"primary_key"`
}
