package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryDeviceByID(id string) (bool, *entity.Device, error) {
	var device entity.Device
	result := database.DbInstance.Select("room_id").Where("device_id = ?", id).Find(&device)

	return result.RowsAffected == 0, &device, result.Error
}
