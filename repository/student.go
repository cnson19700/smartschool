package repository

import (
	"github.com/pkg/errors"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryStudentBySID(sid string) (*entity.Student, bool, error) {
	var student entity.Student
	result := database.DbInstance.Where("student_id = ?", sid).First(&student)

	return &student, result.RowsAffected == 0,  result.Error
}

func QueryAllStudents() ([]*entity.Student, error) {
	students := []*entity.Student{}
	err := database.DbInstance.Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, err
}

func QueryCheckinHistoryWithSIdAndStatus(id int, status string) ([]entity.Attendance, error) {
	var stat []entity.Attendance
	err := database.DbInstance.Where("student_id = ? AND check_in_status = ?", id, status).Preload("Scheduler").Preload("Scheduler.Course").Find(&stat).Error

	if err != nil {
		return nil, err
	}

	result := append([]entity.Attendance{}, stat...)

	return result, nil
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
