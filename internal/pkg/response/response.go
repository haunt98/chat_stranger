package response

var Codes = map[int]string{
	200: "Fetch all OK",
	201: "Found user",
	202: "Update info OK",
	203: "Update password OK",
	204: "Delete OK",
	205: "Register OK",
	206: "Login OK",
	207: "New room OK",
	208: "Next room OK",
	209: "Get empty room OK",
	210: "Join room OK",
	211: "Leave room OK",
	212: "Receive message OK",
	213: "Send message OK",
	214: "Next empty room ok",
	400: "JSON body is wrong",
	401: "id param is wrong",
	402: "Fetch all failed",
	403: "User not exist",
	404: "Register name is already used",
	405: "Register name or password not correct",
	406: "Decode token failed",
	407: "Auth token not found",
	408: "Auth token not valid",
	409: "Role wrong",
	410: "Room is full",
	411: "No room id",
	412: "Register name is missing",
	413: "Password is missing",
	414: "Full name is missing",
	415: "Get empty room failed",
	416: "Join room failed",
	417: "Leave room failed",
	418: "User not in room",
	419: "Receive message failed",
	420: "Send message failed",
	421: "User already in room",
	422: "Token userid mismatch message from user id",
	423: "Next empty room failed",
	500: "Create token failed",
	501: "Get id failed",
}

func Make(code int) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = code
	res["message"] = Codes[code]
	return res
}
