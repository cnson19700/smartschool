package repository

import (
	"time"

	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QuerySemesterByFacultyTime(faculty_id uint, checkTime time.Time) (*entity.Semester, bool, error) {
	var Semester entity.Semester
	result := database.DbInstance.Select("id").Where("faculty_id = ? AND end_time > ? AND start_time <= ?", faculty_id, checkTime, checkTime).Limit(1).Find(&Semester)

	return &Semester, result.RowsAffected == 0,  result.Error
}