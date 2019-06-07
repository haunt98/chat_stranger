package dtos

type UserRequest struct {
	RegName   string `json:"regname"`
	Password  string `json:"password"`
	FullName  string `json:"fullname"`
	Gender    string `json:"gender"`
	BirthYear int    `json:"birthyear"`
	Introduce string `json:"introduce"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	FullName  string `json:"fullname"`
	Gender    string `json:"gender"`
	BirthYear int    `json:"birthyear"`
	Introduce string `json:"introduce"`
}
