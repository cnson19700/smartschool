package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryComplainFormIDByUser(request_user_id, receive_user_id, schedule_id uint) (entity.ComplainForm, bool, error) {
	var form entity.ComplainForm
	result := database.DbInstance.Select("id").Where("request_user_id = ? AND receive_user_id = ? AND schedule_id = ? AND deleted_at IS NULL", request_user_id, receive_user_id, schedule_id).Find(&form)

	return form, result.RowsAffected == 0, result.Error
}

func QueryExistComplainFormIDByUser(request_user_id, schedule_id uint) (bool, error) {
	var form entity.ComplainForm
	result := database.DbInstance.Select("id").Where("request_user_id = ? AND schedule_id = ? AND deleted_at IS NULL", request_user_id, schedule_id).Find(&form)

	return result.RowsAffected == 0, result.Error
}
