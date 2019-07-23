package valid

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

func checkRegisterName(name string) error {
	if name == "" {
		return fmt.Errorf("tên đăng nhập bị thiếu")
	}
	return nil
}

func checkPassword(password string) error {
	if password == "" {
		return fmt.Errorf("mật khẩu bị thiếu")
	}
	return nil
}

func checkShowName(name string) error {
	if name == "" {
		return fmt.Errorf("tên hiển thị bị thiếu")
	}
	return nil
}

func checkGender(gender string) error {
	if gender == "" {
		return fmt.Errorf("giới tính bị thiếu")
	}
	return nil
}

func checkBirthYear(year int) error {
	if year <= 0 {
		return fmt.Errorf("năm sinh không thể là số âm")
	}
	if year > time.Now().Year() {
		return fmt.Errorf("năm sinh không thể lớn hơn năm hiện tại")
	}
	return nil
}

func CheckSignUpSubmit(showName, regName, password string) error {
	if err := checkRegisterName(regName); err != nil {
		return errors.Wrap(err, "check sign up submit failed")
	}
	if err := checkPassword(password); err != nil {
		return errors.Wrap(err, "check sign up submit failed")
	}
	if err := checkShowName(showName); err != nil {
		return errors.Wrap(err, "check sign up submit failed")
	}
	return nil
}

func CheckLogInSubmit(regName, password string) error {
	if err := checkRegisterName(regName); err != nil {
		return errors.Wrap(err, "check log in submit failed")
	}
	if err := checkPassword(password); err != nil {
		return errors.Wrap(err, "check log in submit failed")
	}
	return nil
}

func CheckUpdateInfoSubmit(showName, gender string, birthYear int) error {
	if err := checkShowName(showName); err != nil {
		return errors.Wrap(err, "check update info submit failed")
	}
	if err := checkGender(gender); err != nil {
		return errors.Wrap(err, "check update info submit failed")
	}
	if err := checkBirthYear(birthYear); err != nil {
		return errors.Wrap(err, "check update info submit failed")
	}
	return nil
}
