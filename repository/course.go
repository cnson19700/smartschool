package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryAllCourses() ([]*entity.Course, error) {
	courses := []*entity.Course{}
	err := database.DbInstance.Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func QueryAllCourses() *[]entity.Course {
	var course []entity.Course
	result := database.DbInstance.Find(&course)
	if result.Error != nil {
		return nil

	}

	return &course
}
