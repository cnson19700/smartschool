package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryCourseByID(id string) (*entity.Course, bool, error) {
	var course entity.Course
	result := database.DbInstance.Select("id").Where("course_id = ?", id).Find(&course)

	return &course, result.RowsAffected == 0, result.Error
}

func QueryAllCourses() (*[]entity.Course, error) {
	var course []entity.Course
	err := database.DbInstance.Find(&course).Error
	if err != nil {
		return nil, err
	}

	return &course, nil
}

func QueryCourseInfoByID(id uint) (*entity.Course, bool, error) {
	var course entity.Course
	result := database.DbInstance.Where("id = ?", id).Find(&course)

	return &course, result.RowsAffected == 0, result.Error
}
func QueryCourseByTeacherID(teacher_id string) ([]*entity.CourseByTeacher, error) {
	courses := []*entity.Course{}
	courses_by_teacher := []*entity.CourseByTeacher{}

	err := database.DbInstance.Where("teacher_id = ?", teacher_id).Find(&courses).Error

	if err != nil {
		return nil, err
	}

	for _, course := range courses {
		course := entity.CourseByTeacher{
			ID:              course.ID,
			CourseID:        course.CourseID,
			TeacherID:       course.TeacherID,
			TeacherRole:     course.TeacherRole,
			Name:            course.Name,
			SemesterID:      course.SemesterID,
			NumberOfStudent: course.NumberOfStudent,
		}
		courses_by_teacher = append(courses_by_teacher, &course)
	}

	return courses_by_teacher, nil
}
