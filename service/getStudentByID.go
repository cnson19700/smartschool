package service

import (
	"github.com/smartschool/database"
	"github.com/smartschool/entity"
)

func GetStudentByID(id int) *entity.Student {
	var student entity.Student
	database.DbInstance.Where("id = ?", id).First(&student)
	if student.ID == 0 {
		return nil
	}

	return &student
}
