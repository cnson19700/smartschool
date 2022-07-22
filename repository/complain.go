package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryComplainFormByUser(request_user_id, receive_user_id, schedule_id uint) (entity.ComplainForm, bool, error) {
	var form entity.ComplainForm
	result := database.DbInstance.Where("request_user_id = ? AND receive_user_id = ? AND schedule_id = ?", request_user_id, receive_user_id, schedule_id).Find(&form)

	return form, result.RowsAffected == 0, result.Error
}

func QueryExistComplainFormByUser(request_user_id, schedule_id uint) (bool, error) {
	var form entity.ComplainForm
	result := database.DbInstance.Select("id").Where("request_user_id = ? AND schedule_id = ? AND deleted_at IS NULL", request_user_id, schedule_id).Find(&form)

	return result.RowsAffected == 0, result.Error
}

func QueryListComplainFormByUserSemester(user_id, semester_id uint) ([]entity.ComplainForm, bool, error) {
	var form_list []entity.ComplainForm
	result := database.DbInstance.Where("request_user_id = ? AND semester_id = ? AND deleted_at IS NULL", user_id, semester_id).Find(&form_list)

	return form_list, result.RowsAffected == 0, result.Error
}

func QueryExistComplainFormByUserForm(request_user_id, form_id uint) (bool, error) {
	var form entity.ComplainForm
	result := database.DbInstance.Select("id").Where("id = ? AND request_user_id = ? AND deleted_at IS NULL", form_id, request_user_id).Find(&form)

	return result.RowsAffected == 0, result.Error
}

func QueryComplainFormByID(form_id uint) (entity.ComplainForm, bool, error) {
	var form entity.ComplainForm
	result := database.DbInstance.Where("id = ? AND deleted_at IS NULL", form_id).Find(&form)

	return form, result.RowsAffected == 0, result.Error
}
