package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryDeviceByID(id string) (*entity.Device, error) {
	var device entity.Device
	err := database.DbInstance.Select("room_id").Where("device_id = ?", id).Find(&device).Error
	return &device, err
}
