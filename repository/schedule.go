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

func QueryListScheduleByCourse(course_id uint) ([]entity.Schedule, bool, error) {
	var queryList []entity.Schedule
	result := database.DbInstance.Where("course_id = ?", course_id).Preload("Room").Find(&queryList)

	//scheduleList := append([]entity.Schedule{}, queryList...)

	return queryList, result.RowsAffected == 0, result.Error
}
