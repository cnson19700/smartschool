package repository

import (
	"github.com/smartschool/database"
)

func QueryUserByCardID(id string) (uint, bool, error) {
	var user_id uint
	result := database.DbInstance.Table("cards").Select("user_id").Where("card_id = ?", id).Find(&user_id)

	return user_id, result.RowsAffected == 0, result.Error
}
