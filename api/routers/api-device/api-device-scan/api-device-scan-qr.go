package api_device_scan

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/smartschool/model"
)

const (
	TimeSlotDuration = 30
	IdxLimit         = 5
)

func ScanQRCodeData(c *gin.Context) {
	var requestData model.ScanQRCodeData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		return
	}
	now := time.Now().Unix()

	scanQRData := requestData
}
