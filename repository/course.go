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
