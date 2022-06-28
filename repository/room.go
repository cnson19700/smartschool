package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryRoomInfo(room_id uint) (*entity.Room, bool, error) {
	var room entity.Room
	result := database.DbInstance.Where("id = ?", room_id).Find(&room)

	return &room, result.RowsAffected == 0, result.Error
}

func QueryRoomByName(name string) uint {
	var room *entity.Room
	database.DbInstance.Where("name = ?", name).Find(&room)

	return room.ID
}
