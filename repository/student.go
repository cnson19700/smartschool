package repository

import (
	"github.com/pkg/errors"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryStudentBySID(sid string) (*entity.Student, error) {
	var student entity.Student
	err := database.DbInstance.Where("student_id = ?", sid).First(&student).Error
	if err != nil {
		return nil, err
	}

	return student, nil
}

func QueryAllStudents() ([]*entity.Student, error) {
	students := []*entity.Student{}
	err := database.DbInstance.Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
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

func QueryStudentByEmail(email string) (*entity.User, error) {
	user := &entity.User{}

	err := database.DbInstance.Where("email = ?", email).
		First(user).Error

	if err != nil {
		return nil, errors.Wrap(err, "get user by email")
	}

	return user, nil
}
