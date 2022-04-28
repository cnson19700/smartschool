package repository

import "github.com/smartschool/database"

func ExistQueryEnrollmentByTeacherCourse(teacher_id uint, course_id uint) (bool, error) {
	var verify uint
	result := database.DbInstance.Table("teacher_courses").Select("id").Where("teacher_id = ? AND course_id = ?", teacher_id, course_id).Find(&verify)

	return result.RowsAffected == 0, result.Error
}