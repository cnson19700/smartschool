package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryCourseByID(id string) *entity.Course {
	var course entity.Course
	database.DbInstance.Select("id").Where("course_id = ?", id).Find(&course)
	if course.ID == 0 {
		return nil
	}
	return &course
}
