package repository

import (
	"time"

	"github.com/smartschool/database"
	"github.com/smartschool/model/entity"
)

func QueryScheduleByRoomTimeCourse(room_id string, time time.Time, course_id string) (*entity.Schedule, error) {
	var schedule entity.Schedule
	err := database.DbInstance.Order("end_time").Select("id", "start_time", "end_time").Where("room_id = ? AND end_time >= ? AND course_id = ?", room_id, time, course_id).Find(&schedule).Error
	if err != nil {
		return nil, err
	}
	return &schedule, err
}
