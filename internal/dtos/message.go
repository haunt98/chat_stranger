package dtos

type MessageRequest struct {
	RoomID int    `json:"roomid"`
	Body   string `json:"body"`
}

type MessageResponse struct {
	RoomID   int    `json:"roomid"`
	FromUser string `json:"fromuser"`
	Body     string `json:"body"`
}

type LatestRequest struct {
	RoomID int `json:"roomid"`
	Latest int `json:"latest"`
}
