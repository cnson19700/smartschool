package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func CreateLogCheckIn(deviceSignal entity.DeviceSignalLog) error {
	err := database.DbInstance.Create(&deviceSignal).Error

	return err
}
