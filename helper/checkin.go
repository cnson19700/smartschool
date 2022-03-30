package helper

import (
	"encoding/base64"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const Prefix string = "11"
const SecretKey string = "Keep read as your greed's lead, Each greed thing is a good news, Did you see the news of yourself, You are now cursed by the hell! - KEDY"
const AcceptRefreshQRSecond uint = 30

func ClassifyCheckinCode(code string) (CheckinType string, Value string, err error) {
	reQR, err := regexp.Compile(`^[a-zA-Z0-9]+:\S+\=$`) //format: <Prefix>:<encodeString>=
	if err != nil {
		return "ERROR", "", err
	}

	reCard, err := regexp.Compile("^[a-zA-Z0-9]+$")
	if err != nil {
		return "ERROR", "", err
	}

	if reQR.Match([]byte(code)) {
		// checkCode := code[(strings.Index((code), ":") + 1):(len(code) - 1)]
		// // checkCode = base64.StdEncoding.EncodeToString([]byte(checkCode)) //this is temp
		// rawDecodedText, err := base64.StdEncoding.DecodeString(checkCode)

		return "QR", code, nil

	}

	if reCard.Match([]byte(code)) {
		return "Card", code, nil
	}

	return "ERROR", "", nil
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

func ParseQR(code string) (string, bool, error) {
	code = code[:(len(code) - 1)]
	checkCodeValues := strings.Split(code, ":")
	if len(checkCodeValues) != 2 {
		return "", false, nil
	}

	if checkCodeValues[0] != Prefix {
		return "", false, nil
	}

	rawDecodedText, err := base64.StdEncoding.DecodeString(checkCodeValues[1])
	if err != nil {
		return "", false, err
	}
	contentValues := strings.Split(string(rawDecodedText), "|")
	if len(contentValues) != 3 {
		return "", false, nil
	}

	userId := contentValues[0]
	requestDateTime, err := StringToTimeUTC(contentValues[1])
	secret := contentValues[2]
	if err != nil {
		return "", false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(secret), []byte(SecretKey))
	if err != nil {
		return "", false, nil
	}

	if diff := time.Since(requestDateTime); diff > (time.Second*time.Duration(AcceptRefreshQRSecond)) || diff < 0 {
		return "", false, nil
	}

	return userId, true, nil
}
