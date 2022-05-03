package repository

import (
	"github.com/smartschool/database"
)

func QueryRoomByDeviceID(id string) (uint, bool, error) {
	var room_id uint
	result := database.DbInstance.Table("devices").Select("room_id").Where("device_id = ?", id).Find(&room_id)

	return room_id, result.RowsAffected == 0, result.Error
}
