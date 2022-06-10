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
		return errors.New("password must have 8 characters")
	}
	if newPass != reNewPass {
		return errors.New("2 password not matches")
	}
	return nil
}
