package valid

func CheckRegisterName(name string) bool {
	if name == "" {
		return false
	}
	return true
}

func CheckPassword(password string) bool {
	if password == "" {
		return false
	}
	return true
}

func CheckPassword2(password, password2 string) bool {
	if password == "" {
		return false
	}
	return password == password2
}
