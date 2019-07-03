package model

type Like struct {
	ID      int `json:"id"`
	UserID  int `json:"user_id"`
	HobbyID int `json:"hobby_id"`
}
