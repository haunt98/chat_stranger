package valid

func CheckRegName(regname string) int {
	if regname == "" {
		return 412
	}

	return 0
}

func CheckPassword(password string) int {
	if password == "" {
		return 413
	}

	return 0
}

func CheckFullName(fullname string) int {
	if fullname == "" {
		return 414
	}

	return 0
}
