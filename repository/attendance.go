package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryAttendanceByStudentSchedule(student_id uint, schedule_id uint) (*entity.Attendance, error) {
	var checkAttend entity.Attendance
	err := database.DbInstance.Select("id").Where("user_id = ? AND schedule_id = ?", student_id, schedule_id).Find(&checkAttend).Error

	return &checkAttend, err
}

func CreateAttendance(attendance entity.Attendance) {
	database.DbInstance.Create(&attendance)
}
