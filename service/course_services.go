package service

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func GetCourseByCourseID(id string) (*entity.Course, *error) {
	var course entity.Course
	result := database.DbInstance.Where("COURSE_ID = ?", id).First(&course)
	if result.Error != nil {
		return nil, &result.Error

	}

	return &course, nil
}

func GetCourses() (*[]entity.Course, *error) {
	var course []entity.Course
	result := database.DbInstance.Find(&course)
	if result.Error != nil {
		return nil, &result.Error

	}

	return &course, nil
}
