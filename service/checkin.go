package service

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"github.com/smartschool/model/dto"
)

// c *gin.Context, deviceSignal dto.DeviceSignal

func CheckIn(deviceSignal dto.DeviceSignal) {
	reQR, errMatchQR := regexp.Compile(`^[a-zA-Z0-9]+:\S+\=$`)
	reCard, errMatchCard := regexp.Compile("^[a-zA-Z0-9]+$")

	if errMatchQR != nil {
		fmt.Println("Regex QR Error")
		panic(errMatchQR)
	}

	if errMatchCard != nil {
		fmt.Println("Regex Card Error")
		panic(errMatchCard)
	}

	if reQR.Match([]byte(deviceSignal.CardId)) {
		checkCode := deviceSignal.CardId[(strings.Index((deviceSignal.CardId), ":") + 1):(len(deviceSignal.CardId) - 1)]
		rawDecodedText, err := base64.StdEncoding.DecodeString(checkCode)

		if err != nil {
			panic(err)
		}

		fmt.Println("QR code received: " + string(rawDecodedText))

	} else if reCard.Match([]byte(deviceSignal.CardId)) {
		fmt.Println("Card code received: " + string(deviceSignal.CardId))

	}
}
