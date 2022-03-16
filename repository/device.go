package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryDeviceByID(id string) *entity.Device {
	var device entity.Device
	database.DbInstance.Select("room_id").Where("device_id = ?", id).Find(&device)
	if device.RoomID == 0 {
		return nil
	}
	return &device
}
