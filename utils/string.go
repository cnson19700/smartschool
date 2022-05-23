package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func LimitStringToInt(s string, defaultLimit int) int {
	if s != "" {
		value, err := strconv.Atoi(s)
		if err != nil {
			return defaultLimit
		}
		if value == 0 {
			return 1
		}
		return value
	}

	return defaultLimit
}

func StringFromBool(input bool) string {
	if input {
		return "True"
	}
	return "False"
}

func Contain(target interface{}, list interface{}) (bool, int) {
	if reflect.TypeOf(list).Kind() == reflect.Slice || reflect.TypeOf(list).Kind() == reflect.Array {
		listvalue := reflect.ValueOf(list)
		for i := 0; i < listvalue.Len(); i++ {
			if target == listvalue.Index(i).Interface() {
				return true, i
			}
		}
	}
	if reflect.TypeOf(target).Kind() == reflect.String && reflect.TypeOf(list).Kind() == reflect.String {
		return strings.Contains(list.(string), target.(string)), strings.Index(list.(string), target.(string))
	}
	return false, -1
}

func ArrStringPToArrValue(array []*string) []string {
	var result []string
	for _, v := range array {
		result = append(result, *v)
	}
	return result
}

func ArrValueToArrStringP(array []string) []*string {
	var result []*string
	for i := range array {
		result = append(result, &array[i])
	}
	return result
}

func GetSecondReaction(sec string, defaultSec int) int {
	secInt, err := strconv.Atoi(sec)

	if err != nil || secInt > 10 {
		return defaultSec
	}
	return secInt
}

func GetScheduleTime(second int) string {
	return "@every " + strconv.Itoa(second) + "s"
}

func GetMonthDateS(t *time.Time) string {
	if t == nil {
		return ""
	}
	month, date := t.Month(), t.Day()
	return fmt.Sprintf("%02d月%02d日", month, date)
}

func GetHourMinuteS(t *time.Time) string {
	if t == nil {
		return ""
	}
	hour, minute := t.Hour(), t.Minute()
	return fmt.Sprintf("%02d:%02d", hour, minute)
}

func GetExpressPlayTokenFromUrl(url string) string {
	if strings.Contains(url, "ExpressPlayToken=") {
		return strings.Split(url, "ExpressPlayToken=")[1]
	}
	return ""
}

func StrPEqualValue(p1, p2 *string) bool {
	return StrFromP(p1) == StrFromP(p2)
}

func StrFromP(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func RemoveTabCharacters(s string) string {
	return strings.ReplaceAll(s, "\t", "")
}

func StringPJoin(str []*string, delimiter string) string {
	stringCombine := ""
	for _, c := range str {
		if c != nil {
			stringCombine += *c + delimiter
		}
	}
	stringCombine = strings.TrimSuffix(stringCombine, delimiter)
	return stringCombine
}
