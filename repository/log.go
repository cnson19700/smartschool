package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func CreateLogCheckIn(deviceSignal entity.DeviceSignalLog) {
	database.DbInstance.Create(&deviceSignal)
}
