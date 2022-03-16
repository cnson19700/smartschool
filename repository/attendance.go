package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryAttendanceByStudentSchedule(student_id uint, schedule_id uint) *entity.Attendance {
	var checkAttend entity.Attendance
	database.DbInstance.Select("id").Where("user_id = ? AND schedule_id = ?", student_id, schedule_id).Find(&checkAttend)
	if checkAttend.ID == 0 {
		return nil
	}
	return &checkAttend
}

func CreateAttendance(attendance entity.Attendance) {
	database.DbInstance.Create(&attendance)
}
