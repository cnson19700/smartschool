package repository

import (
	"time"

	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryAttendanceByStudentSchedule(student_id uint, schedule_id uint) (*entity.Attendance, bool, error) {
	var checkAttend entity.Attendance
	result := database.DbInstance.Select("id").Where("user_id = ? AND schedule_id = ?", student_id, schedule_id).Find(&checkAttend)

	return &checkAttend, result.RowsAffected == 0, result.Error
}

func CreateAttendance(attendance entity.Attendance) error {
	err := database.DbInstance.Create(&attendance).Error

	return err
}

func QueryListAttendanceByUserSchedule(user_id uint, schedule_id_list []uint) ([]entity.Attendance, bool, error) {
	var queryList []entity.Attendance
	result := database.DbInstance.Where("user_id = ? AND schedule_id IN ?", user_id, schedule_id_list).Find(&queryList)

	//attendanceList := append([]entity.Attendance{}, queryList...)

	return queryList, result.RowsAffected == 0, result.Error
}

func QueryListAttendanceInDayByUser(user_id uint, start time.Time, end time.Time) ([]entity.Attendance, bool, error) {
	var queryList []entity.Attendance
	result := database.DbInstance.Where("user_id = ? AND (checkin_time BETWEEN ? AND ?)", user_id, start, end).Preload("Schedule.Room").Preload("Schedule.Course").Find(&queryList)

	return queryList, result.RowsAffected == 0, result.Error
}

func CountAttendanceOfSchedule(user_id uint, schedule_id_list []uint) (int64, error) {
	var c int64
	result := database.DbInstance.Table("attendances").Select("id").Where("user_id = ? AND schedule_id IN ?", user_id, schedule_id_list).Count(&c)

	return c, result.Error
}
