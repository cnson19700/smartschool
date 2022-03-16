package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryEnrollmentByStudentCourse(student_id string, course_id string) (*entity.StudentCourseEnrollment, error) {
	var verify entity.StudentCourseEnrollment
	err := database.DbInstance.Select("id").Where("student_id = ? AND course_id = ?", student_id, course_id).Find(&verify).Error
	if err != nil {
		return nil, err
	}
	return &verify, nil
}
