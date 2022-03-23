package repository

import (
	"time"

	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryScheduleByRoomTime(room_id uint, checkinTime time.Time) (bool, *entity.Schedule, error) {
	var schedule entity.Schedule
	result := database.DbInstance.Order("end_time").Select("id", "course_id", "start_time", "end_time").Where("room_id = ? AND end_time >= ?", room_id, checkinTime).Find(&schedule)

	return result.RowsAffected == 0, &schedule, result.Error
}

func QueryScheduleByRoomTimeCourse(room_id uint, time time.Time, course_id uint) (bool, *entity.Schedule, error) {
	var schedule entity.Schedule
	result := database.DbInstance.Order("end_time").Select("id", "start_time", "end_time").Where("room_id = ? AND end_time >= ? AND course_id = ?", room_id, time, course_id).Find(&schedule)

	return result.RowsAffected == 0, &schedule, result.Error

}
