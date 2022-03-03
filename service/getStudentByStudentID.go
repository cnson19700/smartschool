package service

import (
	"github.com/smartschool/database"
	"github.com/smartschool/entity"
)

func GetStudentByID(id string) *entity.Student {
	var student entity.Student
	database.DbInstance.Where("student_id = ?", id).First(&student)
	if student.ID == 0 {
		return nil
	}

	return &student
}
