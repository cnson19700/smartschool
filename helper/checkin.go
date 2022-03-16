package helper

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func ClassifyCheckinCode(code string) (CheckinType string, Value string) {
	reQR, errMatchQR := regexp.Compile(`^[a-zA-Z0-9]+:\S+\=$`) //format: <Prefix>:<encodeString>=
	reCard, errMatchCard := regexp.Compile("^[a-zA-Z0-9]+$")

	if errMatchQR != nil {
		panic(errMatchQR)
	}

	if errMatchCard != nil {
		panic(errMatchCard)
	}

	if reQR.Match([]byte(code)) {
		checkCode := code[(strings.Index((code), ":") + 1):(len(code) - 1)]
		checkCode = base64.StdEncoding.EncodeToString([]byte(checkCode)) //this is temp
		rawDecodedText, err := base64.StdEncoding.DecodeString(checkCode)
		if err != nil {
			fmt.Println("error:", err)
		}

		return "QR", string(rawDecodedText)

	} else if reCard.Match([]byte(code)) {
		return "Card", code
	}

	return "", ""
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
