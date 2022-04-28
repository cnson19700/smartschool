package helper

import (
	"encoding/base64"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/smartschool/lib/constant"
	"golang.org/x/crypto/bcrypt"
)

func ClassifyCheckinCode(code string) (CheckinType string, err error) {
	reQR, err := regexp.Compile(`^[a-zA-Z0-9]+:\S+\=$`) //format: <Prefix>:<encodeString>=
	if err != nil {
		return "ERROR", err
	}

	reCard, err := regexp.Compile("^[a-zA-Z0-9]+$")
	if err != nil {
		return "ERROR", err
	}

	if reQR.Match([]byte(code)) {
		// checkCode := code[(strings.Index((code), ":") + 1):(len(code) - 1)]
		// // checkCode = base64.StdEncoding.EncodeToString([]byte(checkCode)) //this is temp
		// rawDecodedText, err := base64.StdEncoding.DecodeString(checkCode)

		return "QR", nil

	}

	if reCard.Match([]byte(code)) {
		return "Card", nil
	}

	return "ERROR", nil
}

func ConvertDeviceTimestampToExact(timestamp int64) time.Time {
	tempTime := time.Unix(timestamp, 0)
	tempTime = tempTime.Add((-1) * time.Hour * 7)
	return tempTime
}

func ParseData(checkinValues string) (string, string) {
	res := strings.Split(checkinValues, "-")
	return res[0], res[1]
}

func CheckValidDifferentTimeEntry(timeEntry time.Time, acceptDuration time.Duration) bool {
	if diff := time.Since(timeEntry); diff >= 0 && diff < acceptDuration {
		return true
	}

	return false
}

func ParseQR(code string, timeEntry time.Time) (uint, bool, error) {
	code = code[:(len(code) - 1)]
	checkCodeValues := strings.Split(code, ":")
	if len(checkCodeValues) != 2 {
		return 0, false, nil
	}

	if checkCodeValues[0] != constant.QRPrefix {
		return 0, false, nil
	}

	rawDecodedText, err := base64.StdEncoding.DecodeString(checkCodeValues[1])
	if err != nil {
		return 0, false, err
	}
	contentValues := strings.Split(string(rawDecodedText), "|")
	if len(contentValues) != 3 {
		return 0, false, nil
	}

	userId_str := contentValues[0]

	userId, err := strconv.ParseUint(userId_str, 10, 64)
	if err != nil {
		return 0, false, err
	}

	requestDateTime, err := StringToTimeUTC(contentValues[1])
	secret := contentValues[2]
	if err != nil {
		return uint(userId), false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(secret), []byte(constant.QRSecretKey+contentValues[1]))
	if err != nil {
		return uint(userId), false, nil
	}

	// if diff := time.Since(requestDateTime); diff > constant.AcceptRefreshQRSecond || diff < 0 {
	// 	return "", false, nil
	// }
	if diff := timeEntry.Sub(requestDateTime); diff > constant.AcceptRefreshQRSecond || diff < 0 {
		return uint(userId), false, nil
	}

	return uint(userId), true, nil
}
