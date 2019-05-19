package handler

func Response(status bool, message string) map[string]interface{} {
	res := make(map[string]interface{})
	res["Status"] = status
	res["Message"] = message
	return res
}
