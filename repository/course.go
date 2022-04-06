package repository

import (
	"github.com/smartschool/database"
	"github.com/smartschool/model/dto"
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

func QueryCourseBasicInfoByID(id uint) (*dto.CourseReportPartElement, bool, error) {
	var course dto.CourseReportPartElement
	result := database.DbInstance.Table("courses").Where("id = ?", id).Find(&course)

	return &course, result.RowsAffected == 0, result.Error
}

func QueryListCourseBySemester(sem_id uint) ([]dto.CourseReportPartElement, bool, error) {
	var queryList []dto.CourseReportPartElement
	result := database.DbInstance.Table("courses").Where("semester_id = ?", sem_id).Find(&queryList)

	return queryList, result.RowsAffected == 0, result.Error
}

func QueryListCourseIDBySemester(sem_id uint) ([]uint, bool, error) {
	var queryList []uint
	result := database.DbInstance.Table("courses").Select("id").Where("semester_id = ?", sem_id).Find(&queryList)

	return queryList, result.RowsAffected == 0, result.Error
}

func QueryListCourseBasicInfoByID(list_id []uint) ([]dto.CourseReportPartElement, bool, error) {
	var queryList []dto.CourseReportPartElement
	result := database.DbInstance.Table("courses").Where("id IN ?", list_id).Find(&queryList)

	return queryList, result.RowsAffected == 0, result.Error
}
