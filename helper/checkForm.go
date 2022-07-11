package helper

import (
	"errors"
	"strconv"
)

func CheckAgeFormat(age int) (bool, string) {
	minAge := 0
	maxAge := 112

	if age < minAge || age > maxAge {
		return false, errors.New("age is not invalid").Error()
	}
	return true, strconv.Itoa(age)
}

func CheckMailFormat(email string) (bool, string) {
	if !ValidEmail(email) {
		return false, errors.New("email format is not invalid").Error()
	}
	return true, email
}

func CompareOldNewPass(oldPass, newPass string) error {
	if oldPass == newPass {
		return errors.New("old and new passwords must be different")
	}
	return nil
}

func ComparePassword(newPass, reNewPass string) error {
	if len(reNewPass) < 8 {
		return errors.New("Mật khẩu phải bao gồm ít nhất 8 ký tự!")
	}
	if newPass != reNewPass {
		return errors.New("Mật khẩu mới và mật khẩu xác nhận không trùng khớp!")
	}
	return nil
}
