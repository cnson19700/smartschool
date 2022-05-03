package repository

import (
	"github.com/smartschool/database"
)

func ExistEnrollmentByStudentCourse(student_id uint, course_id uint) (bool, error) {
	var verify uint
	result := database.DbInstance.Table("student_course_enrollments").Select("id").Where("student_id = ? AND course_id = ?", student_id, course_id).Find(&verify)

	return result.RowsAffected == 0, result.Error
}

func QueryEnrollmentByListCourse(user_id uint, list_course []uint) ([]uint, bool, error) {
	var queryList []uint
	result := database.DbInstance.Table("student_course_enrollments").Select("course_id").Where("student_id = ? AND course_id IN ?", user_id, list_course).Find(&queryList)

	return queryList, result.RowsAffected == 0, result.Error
}
