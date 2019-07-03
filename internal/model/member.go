package model

type Member struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	RoomID int `json:"room_id"`
}
