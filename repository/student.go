package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryStudentBySID(sid string) *entity.Student {
	var student entity.Student
	database.DbInstance.Where("student_id = ?", sid).First(&student)
	if student.ID == 0 {
		return nil
	}

	return &student
}

func QueryCheckinHistoryWithSIdAndStatus(id int, status string) []entity.Attendance {
	var stat []entity.Attendance
	database.DbInstance.Where("student_id = ? AND check_in_status = ?", id, status).Preload("Scheduler").Preload("Scheduler.Course").Find(&stat)

	if len(stat) == 0 {
		return nil
	}

	result := append([]entity.Attendance{}, stat...)

	return result
}
