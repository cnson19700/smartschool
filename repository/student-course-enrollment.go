package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func ExistEnrollmentByStudentCourse(student_id uint, course_id uint) (bool, error) {
	var verify uint
	result := database.DbInstance.Table("student_course_enrollments").Select("id").Where("student_id = ? AND course_id = ?", student_id, course_id).Find(&verify)

	return result.RowsAffected == 0, result.Error
}

func QueryEnrollmentByListCourse(user_id uint, list_course []uint) ([]uint, bool, error) {
	var queryList []uint
	result := database.DbInstance.Table("student_course_enrollments").Select("course_id").Where("student_id = ? AND course_id IN ? AND deleted_at IS NULL", user_id, list_course).Find(&queryList)

	return queryList, result.RowsAffected == 0, result.Error
}

func QueryFirstStudentInCourseEnrollment(course_id int) (int, error) {
	var student_course_enrollment *entity.StudentCourseEnrollment

	err := database.DbInstance.Where("course_id = ?", course_id).First(&student_course_enrollment).Error

	if err != nil {
		return 0, err
	}
	return int(student_course_enrollment.StudentID), nil
}
