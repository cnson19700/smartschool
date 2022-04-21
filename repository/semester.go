package repository

import (
	"time"

	"github.com/smartschool/database"
	"github.com/smartschool/model/dto"
)

func QuerySemesterByFacultyTime(faculty_id uint, checkTime time.Time) (uint, bool, error) {
	var SemesterID uint
	result := database.DbInstance.Table("semesters").Select("id").Where("faculty_id = ? AND end_time > ? AND start_time <= ?", faculty_id, checkTime, checkTime).Limit(1).Find(&SemesterID)

	return SemesterID, result.RowsAffected == 0, result.Error
}

func QuerySemesterByFaculty(faculty_id uint) ([]dto.SemesterListElement, bool, error) {
	var queryList []dto.SemesterListElement
	result := database.DbInstance.Table("semesters").Where("faculty_id = ?", faculty_id).Find(&queryList)

	return queryList, result.RowsAffected == 0, result.Error
}
