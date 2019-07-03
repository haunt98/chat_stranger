package model

type Member struct {
	UserID int `json:"user_id" gorm:"primary_key"`
	RoomID int `json:"room_id" gorm:"primary_key"`
}
