package valid

func CheckRegisterName(name string) (bool, string) {
	if name == "" {
		return false, "Tên đăng nhập bị thiếu"
	}
	return true, ""
}

func CheckPassword(password string) (bool, string) {
	if password == "" {
		return false, "Mật khẩu bị thiếu"
	}
	return true, ""
}

func CheckFullName(name string) (bool, string) {
	if name == "" {
		return false, "Tên hiển thị bị thiếu"
	}
	return true, ""
}
