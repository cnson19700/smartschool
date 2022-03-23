package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryDeviceByID(id string) (*entity.Device, bool, error) {
	var device entity.Device
	result := database.DbInstance.Select("room_id").Where("device_id = ?", id).Find(&device)

	return &device, result.RowsAffected == 0, result.Error
}
