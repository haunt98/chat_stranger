package dtos

type MessageRequest struct {
	RoomID  int    `json:"roomid"`
	Body string `json:"body"`
}

type MessageResponse struct {
	Body   string `json:"body"`
}
