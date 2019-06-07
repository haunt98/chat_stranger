package dtos

type AdminRequest struct {
	RegName  string `json:"regname"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
}

type AdminResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
}
