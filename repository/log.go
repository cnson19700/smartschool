package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func LogCheckIn(deviceSignal entity.DeviceSignalLog) {
	database.DbInstance.Create(&deviceSignal)
}
