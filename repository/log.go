package repository

import (
	"time"

	"github.com/smartschool/database"
	"github.com/smartschool/model/dto"
	"github.com/smartschool/model/entity"
)

func LogCheckIn(deviceSignal dto.DeviceSignal, status string) {
	log := entity.DeviceSignalLog{CardId: deviceSignal.CardId, CompanyTokenKey: deviceSignal.CompanyTokenKey, Status: status, Timestamp: time.Unix(deviceSignal.Timestamp, 0).Add(-7 * time.Hour)}
	database.DbInstance.Create(&log)
}
