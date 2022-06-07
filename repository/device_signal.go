package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryUserIDByDeviceSignalID(id string) (*entity.DeviceSignalLog, bool, error) {
	var device_signal_logs *entity.DeviceSignalLog
	result := database.DbInstance.Table("device_signal_logs").Where("id = ?", id).Find(&device_signal_logs)

	return device_signal_logs, result.RowsAffected == 0, result.Error
}
