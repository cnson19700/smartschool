package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryAttendanceByStudentSchedule(student_id string, schedule_id uint) (*entity.Attendance, error) {
	var checkAttend entity.Attendance
	err := database.DbInstance.Select("id").Where("student_id = ? AND scheduler_id = ?", student_id, schedule_id).Find(&checkAttend).Error
	if err != nil {
		return nil, err
	}
	return &checkAttend, nil
}
