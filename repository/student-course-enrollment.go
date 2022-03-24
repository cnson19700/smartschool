package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryEnrollmentByStudentCourse(student_id uint, course_id uint) (*entity.StudentCourseEnrollment, bool, error) {
	var verify entity.StudentCourseEnrollment
	result := database.DbInstance.Select("id").Where("student_id = ? AND course_id = ?", student_id, course_id).Find(&verify)

	return &verify, result.RowsAffected == 0, result.Error
}
