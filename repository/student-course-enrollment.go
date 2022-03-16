package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryEnrollmentByStudentCourse(student_id uint, course_id uint) *entity.StudentCourseEnrollment {
	var verify entity.StudentCourseEnrollment
	database.DbInstance.Select("id").Where("student_id = ? AND course_id = ?", student_id, course_id).Find(&verify)
	if verify.ID == 0 {
		return nil
	}
	return &verify
}
