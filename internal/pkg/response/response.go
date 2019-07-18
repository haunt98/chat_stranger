package response

var Codes = map[int]string{
	100: "Sign up Ok",
	101: "Tên đăng nhập đã tồn tại",
	102: "Failed to bind json when sign up",
	103: "Failed to create token when sign up",
	200: "Log in OK",
	201: "Tên đăng nhập hay mật khẩu sai",
	202: "Failed to bind json when log in",
	203: "Failed to create token when log in",
	300: "Get info OK",
	301: "Không tồn tại tài khoản này",
	400: "Find room OK",
	401: "Failed to find room",
	500: "Join room OK",
	501: "Failed to join room",
	502: "Failed to convert roomID to int",
	600: "Leave room OK",
	601: "Failed to leave room",
	700: "Send message OK",
	701: "Failed to send message",
	702: "Failed to bind json when send message",
	800: "Receive message OK",
	801: "Failed to receive message",
	802: "Failed to query from time",
	900: "User is free",
	901: "User is already joined",
	110: "Count member OK",
	111: "Failed to count member",
	120: "Cập nhập thông tin OK",
	121: "Thất bại khi cập nhập thông tin",
	122: "Failed to bind json when update info",
	130: "Update password OK",
	131: "Failed to update password",
	132: "Failed to bind json when update password",
	141: "Valid check failed",
	999: "You are not allowed to do this",
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

func CreateWithMessage(code int, message string) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = code
	res["message"] = message
	return res
}
