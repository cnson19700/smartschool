package helper

import (
	"encoding/base64"
	"regexp"
	"strings"
)

func ClassifyCheckinCode(code string) (CheckinType string, Value string) {
	reQR, errMatchQR := regexp.Compile(`^[a-zA-Z0-9]+:\S+\=$`)
	reCard, errMatchCard := regexp.Compile("^[a-zA-Z0-9]+$")

	if errMatchQR != nil {
		panic(errMatchQR)
	}

	if errMatchCard != nil {
		panic(errMatchCard)
	}

	if reQR.Match([]byte(code)) {
		checkCode := code[(strings.Index((code), ":") + 1):(len(code) - 1)]
		rawDecodedText, err := base64.StdEncoding.DecodeString(checkCode)

		if err != nil {
			panic(err)
		}

		return "QR", string(rawDecodedText)

	} else if reCard.Match([]byte(code)) {
		return "Card", code
	}

	return "", ""
}
