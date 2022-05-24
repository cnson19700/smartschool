package dto

import (
	"time"

	"github.com/smartschool/lib/constant"
)

type RequestChangeAttendanceForm struct {
	ScheduleID    uint                   `json:"schedule_id"`
	CourseName    string                 `json:"course_name"`
	StartTime     time.Time              `json:"start_time"`
	EndTime       time.Time              `json:"end_time"`
	Room          string                 `json:"room"`
	CurrentStatus constant.CheckInStatus `json:"current_status"`
	CheckInTime   *time.Time             `json:"checkin_time"`
}

type TeacherList struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
