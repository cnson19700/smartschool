package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func ExistEnrollmentByTeacherCourse(teacher_id uint, course_id uint) (bool, error) {
	var verify uint
	result := database.DbInstance.Table("teacher_courses").Select("id").Where("teacher_id = ? AND course_id = ?", teacher_id, course_id).Find(&verify)

	return result.RowsAffected == 0, result.Error
}

func DeleteTeacherCourseByListCourseID(list_course_id []uint) error {
	var registers []entity.TeacherCourse
	err := database.DbInstance.Where("course_id IN ?", list_course_id).Delete(&registers).Error

	return err
}

func QueryEnrollmentOfTeacherInCourse(course_id uint) ([]uint, bool, error) {
	var teacherID_list []uint
	result := database.DbInstance.Table("teacher_courses").Select("teacher_id").Where("course_id = ? AND deleted_at IS NULL", course_id).Find(&teacherID_list)

	return teacherID_list, result.RowsAffected == 0, result.Error
}
