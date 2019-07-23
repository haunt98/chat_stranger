package response

var Codes = map[int]string{
	100: "Sign up Ok",
	101: "Tên đăng nhập đã tồn tại",
	102: "Sign up submit failed to parse json",
	103: "Sign up submit is invalid",
	200: "Log in OK",
	201: "Tên đăng nhập hay mật khẩu sai",
	202: "Log in submit failed to parse json",
	203: "Log in submit invalid",
	300: "Get info OK",
	301: "Get info failed",
	400: "Find room OK",
	401: "Find any room failed",
	402: "Find next room failed",
	403: "Find same gender room failed",
	404: "Find same birth year room failed",
	405: "Query status room is invalid",
	500: "Join room OK",
	501: "Join room failed",
	502: "Failed to convert roomID to int",
	600: "Leave room OK",
	601: "Failed to leave room",
	700: "Send message OK",
	701: "Send message failed",
	702: "Send message failed to parse json",
	800: "Receive message OK",
	801: "Receive message failed",
	802: "Receive message failed to query from time",
	900: "User is free",
	901: "User is already joined",
	902: "Check user is free failed",
	110: "Count member OK",
	111: "Count member failed",
	120: "Update info OK",
	121: "Update info failed",
	122: "Update info submit failed to parse json",
	123: "Update info submit is invalid",
	141: "Check valid failed",
	999: "You are not allowed to do this",
}

type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Create(code int) Response {
	return Response{
		Code:    code,
		Message: Codes[code],
	}
}

func CreateWithData(code int, data interface{}) Response {
	return Response{
		Code:    code,
		Message: Codes[code],
		Data:    data,
	}
}

func CreateWithMessage(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
	}
}
