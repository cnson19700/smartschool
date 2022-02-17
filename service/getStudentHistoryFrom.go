package service

import (
	"github.com/smartschool/database"
	"github.com/smartschool/entity"
)

func GetStudentHistoryFrom(id int, status string) *[]entity.Attendance {
	var stat []entity.Attendance
	//DbInstance.Preload("Scheduler", "schedulers.course_id = ?", 2).Where("student_id = ? AND check_in_status = ?", 1, "Late").Find(&stat)
	database.DbInstance.Where("student_id = ? AND check_in_status = ?", id, status).Preload("Scheduler").Preload("Scheduler.Course").Find(&stat)

	// var result []entity.Attendance
	// for i := 0; i < len(stat); i++ {
	// 	result = append(result, entity.Attendance{ID: stat[i].ID, StudentID: stat[i].StudentID, CheckInTime: stat[i].CheckInTime, CheckInStatus: stat[i].CheckInStatus})
	// }
	result := append([]entity.Attendance{}, stat...)

	return &result
}
