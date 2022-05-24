package repository

import "github.com/smartschool/database"

func QueryListFormIDByUser(request_user_id, receive_user_id, semester_id uint) ([]uint, error) {
	var form_id_list []uint
	result := database.DbInstance.Table("attendance_forms").Select("id").Where("request_user_id = ? AND receive_user_id = ? AND semester_id = ?", request_user_id, receive_user_id, semester_id).Find(&form_id_list)

	return form_id_list, result.Error
}

func ExistFormSchedule(form_id_list []uint, schedule_id uint) (bool, error) {
	var id uint
	result := database.DbInstance.Table("form_details").Select("id").Where("attendance_form_id IN ? AND schedule_id = ?", form_id_list, schedule_id).Find(&id)

	return result.RowsAffected == 0, result.Error
}
