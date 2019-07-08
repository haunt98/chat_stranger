package response

var Codes = map[int]string{
	1:    "Đăng ký OK",
	10:   "Tên đăng nhập đã tồn tại",
	11:   "Failed to bind json when sign up",
	12:   "Failed to create token when sign up",
	2:    "Đăng nhập OK",
	20:   "Tên đăng nhập hay mật khẩu sai",
	21:   "Failed to bind json when log in",
	22:   "Failed to create token when log in",
	3:    "Lấy thông tin cá nhân OK",
	30:   "Không tồn tại tài khoản này",
	4:    "Find room OK",
	40:   "Failed to find room",
	5:    "Join room OK",
	50:   "Failed to join room",
	51:   "Failed to convert roomID to int",
	6:    "Leave room OK",
	60:   "Faild to leave room",
	7:    "Send message OK",
	70:   "Failed to send message",
	71:   "Failed to bind json when send message",
	8:    "Receive message OK",
	80:   "Failed to receive message",
	81:   "Failed to query from time",
	9:    "User is free",
	90:   "User already joined",
	901:  "Count member OK",
	9010: "Failed to count member",
	999:  "You are not allowed to do this",
}

func Create(code int) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = code
	res["message"] = Codes[code]
	return res
}

func CreateWithData(code int, data interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = code
	res["message"] = Codes[code]
	res["data"] = data
	return res
}
