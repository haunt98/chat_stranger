package dtos

import "github.com/1612180/chat_stranger/internal/pkg/valid"

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

func (u *UserRequest) Check() int {
	if c:= valid.CheckRegName(u.RegName); c!= 0 {
		return c
	}

	if c := valid.CheckPassword(u.Password); c != 0 {
		return c
	}

	if c := valid.CheckFullName(u.FullName); c != 0 {
		return c
	}

	return 0
}
