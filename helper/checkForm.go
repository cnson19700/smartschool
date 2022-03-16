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

func ComparePassword(oldPass, newPass string) error {
	if len(newPass) < 8 {
		return errors.New("password must have 8 characters")
	}
	if oldPass != newPass {
		return errors.New("2 password not matches")
	}
	return nil
}
