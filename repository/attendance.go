package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryAttendanceByStudentSchedule(student_id uint, schedule_id uint) (*entity.Attendance, bool, error) {
	var checkAttend entity.Attendance
	result := database.DbInstance.Select("id").Where("user_id = ? AND schedule_id = ?", student_id, schedule_id).Find(&checkAttend)

	return &checkAttend, result.RowsAffected == 0, result.Error
}

func CreateAttendance(attendance entity.Attendance) error {
	err := database.DbInstance.Create(&attendance).Error

	return err
}
