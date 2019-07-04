package response

var Codes = map[int]string{
	1:   "OK",
	999: "Vuon hong ngay xua da ua tan",
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
