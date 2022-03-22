package repository

import (
	"time"

	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryScheduleByRoomTime(room_id uint, checkinTime time.Time) (*entity.Schedule, error) {
	var schedule entity.Schedule
	err := database.DbInstance.Order("end_time").Select("id", "course_id", "start_time", "end_time").Where("room_id = ? AND end_time >= ?", room_id, checkinTime).Find(&schedule).Error

	return &schedule, err
}

func QueryScheduleByRoomTimeCourse(room_id uint, time time.Time, course_id uint) (*entity.Schedule, error) {
	var schedule entity.Schedule
	err := database.DbInstance.Order("end_time").Select("id", "start_time", "end_time").Where("room_id = ? AND end_time >= ? AND course_id = ?", room_id, time, course_id).Find(&schedule).Error

	return &schedule, err
}
