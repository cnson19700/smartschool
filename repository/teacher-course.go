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