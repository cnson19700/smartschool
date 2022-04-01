package repository

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryStudentBySID(sid string) (*entity.Student, bool, error) {
	var student entity.Student
	result := database.DbInstance.Where("student_id = ?", sid).Limit(1).Find(&student)

	return &student, result.RowsAffected == 0, result.Error
}

func QueryStudentByID(id string) (*entity.Student, error) {
	var student entity.Student
	err := database.DbInstance.Where("id = ?", id).First(&student).Error
	if err != nil {
		return nil, err
	}

	return &student, nil
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

func QueryStudentsByName(student_name string) ([]uint, error) {
	student_ids := []uint{}
	student_name = strings.ToLower(student_name)

	err := database.DbInstance.Table("students").
		Select("students.id").
		Joins("JOIN users ON users.id = students.id").
		Where("LOWER(CONCAT(users.first_name,users.last_name)) LIKE ?", "%"+student_name+"%").
		Scan(&student_ids).Error
	if err != nil {
		return nil, errors.Wrap(err, "get user by name")
	}
	return student_ids, nil
}
