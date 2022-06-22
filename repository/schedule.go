package repository

import (
	"time"

	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryScheduleByRoomTime(room_id uint, checkinTime time.Time) (*entity.Schedule, bool, error) {
	var schedule entity.Schedule
	result := database.DbInstance.Order("end_time").Select("id", "course_id", "start_time", "end_time").Where("room_id = ? AND end_time > ? AND start_time <= ?", room_id, checkinTime, checkinTime).Limit(1).Find(&schedule)

	return &schedule, result.RowsAffected == 0, result.Error
}

func QueryScheduleByRoomTimeCourse(room_id uint, time time.Time, course_id uint) (*entity.Schedule, bool, error) {
	var schedule entity.Schedule
	result := database.DbInstance.Order("end_time").Select("id", "start_time", "end_time").Where("room_id = ? AND end_time >= ? AND course_id = ?", room_id, time, course_id).Find(&schedule)

	return &schedule, result.RowsAffected == 0, result.Error

}

func QueryFullListScheduleByCourse(course_id uint) ([]entity.Schedule, bool, error) {
	var queryList []entity.Schedule
	result := database.DbInstance.Where("course_id = ?", course_id).Preload("Room").Find(&queryList)

	//scheduleList := append([]entity.Schedule{}, queryList...)

	return queryList, result.RowsAffected == 0, result.Error
}

func QueryListScheduleByListCourseTime(course_id_list []uint, start time.Time, end time.Time) ([]entity.Schedule, bool, error) {
	var queryList []entity.Schedule
	result := database.DbInstance.Order("end_time").Where("course_id IN ? AND start_time >= ? AND end_time <= ?", course_id_list, start, end).Preload("Room").Preload("Course").Find(&queryList)

	return queryList, result.RowsAffected == 0, result.Error
}

func CountScheduleOfFullCourse(course_id uint) (int64, error) {
	var c int64
	result := database.DbInstance.Table("schedules").Select("id").Where("course_id = ?", course_id).Count(&c)

	return c, result.Error
}

func QueryCurrentScheduleIDOfCourse(course_id uint, current time.Time) ([]uint, bool, error) {
	var queryList []uint
	result := database.DbInstance.Table("schedules").Select("id").Where("course_id = ? AND end_time <= ?", course_id, current).Find(&queryList)

	return queryList, result.RowsAffected == 0, result.Error
}

func QueryListScheduleByCourse(course_id uint, current time.Time) ([]entity.Schedule, bool, error) {
	var queryList []entity.Schedule
	result := database.DbInstance.Order("end_time").Where("course_id = ? AND end_time <= ?", course_id, current).Preload("Room").Find(&queryList)

	return queryList, result.RowsAffected == 0, result.Error
}

func QueryScheduleByCourseID(courseID uint) (*entity.Schedule, error) {
	var schedule entity.Schedule
	result := database.DbInstance.Where("course_id = ?", courseID).Find(&schedule)

	return &schedule, result.Error
}

func QueryScheduleByID(ID string) (*entity.Schedule, error) {
	var schedule entity.Schedule
	result := database.DbInstance.Where("id = ?", ID).Find(&schedule)

	return &schedule, result.Error
}
