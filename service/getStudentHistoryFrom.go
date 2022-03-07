package service

import (
	"github.com/smartschool/database"
	"github.com/smartschool/entity"
)

func GetStudentHistoryFrom(id int, status string) *[]entity.Attendance {
	var stat []entity.Attendance

	database.DbInstance.Where("student_id = ? AND check_in_status = ?", id, status).Preload("Scheduler").Preload("Scheduler.Course").Find(&stat)

	result := append([]entity.Attendance{}, stat...)

	return &result
}
