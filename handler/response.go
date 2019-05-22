package handler

func Response(status bool, message string) map[string]interface{} {
	res := make(map[string]interface{})
	res["status"] = status
	res["message"] = message
	return res
}
