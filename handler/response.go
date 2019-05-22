package handler

var ResponseCode = map[int]string{
	200: "Fetch all OK",
	201: "Found user",
	202: "Update info OK",
	203: "Update password OK",
	204: "Delete OK",
	205: "Register OK",
	206: "Login OK",
	400: "JSON body is wrong",
	401: "id param is wrong",
	402: "Fetch all failed",
	403: "User not exist",
	404: "Username is already used",
	405: "Username or password not correct",
	406: "Decode token failed",
	407: "Token not found",
	408: "Token not valid",
	409: "Role wrong",
	500: "Create token failed",
	501: "Get ID failed",
}

func Response(code int) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = code
	res["message"] = ResponseCode[code]
	return res
}
