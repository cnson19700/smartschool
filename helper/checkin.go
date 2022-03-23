package helper

import (
	"encoding/base64"
	"regexp"
	"strings"
	"time"
)

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
		checkCode := code[(strings.Index((code), ":") + 1):(len(code) - 1)]
		// checkCode = base64.StdEncoding.EncodeToString([]byte(checkCode)) //this is temp
		rawDecodedText, err := base64.StdEncoding.DecodeString(checkCode)

		if err != nil {
			return "ERROR", "", err
		}
		return "QR", string(rawDecodedText), nil

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
