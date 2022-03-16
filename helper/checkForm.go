package helper

import (
	"errors"
	"strconv"
	"strings"
)

func CheckFormatValue(formAtributeName string, value string) (bool, string) {
	value = RemoveDoubleSpace(value)
	if value == "" {
		return false, errors.New("auth request is not invalid").Error()
	}
	switch formAtributeName {
	case "email":
		if !ValidEmail(value) {
			return false, errors.New("email format is not invalid").Error()
		}
		return true, value
	case "age":
		age, err := strconv.Atoi(value)
		minAge := 0
		maxAge := 112

		if err != nil || age < minAge || age > maxAge {
			return false, errors.New("age is not invalid").Error()
		}

		return true, value
	case "password":
		str := FormatText(value, true, true)
		if str == "" || strings.Contains(str, " ") {
			return false, formAtributeName + errors.New("password format is not invalid").Error()
		}
		// } else if len(str) <= 8 {
		// 	return false, appErr.AuthMsg.Min8Character
		// }
		return true, str
	default:
		return true, value
	}
}
