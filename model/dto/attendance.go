package dto

import (
	"time"

	"github.com/smartschool/lib/constant"
)

type AttendanceListElement struct {
	ScheduleID    uint                   `json:"schedule_id"`
	StartTime     time.Time              `json:"start_time"`
	EndTime       time.Time              `json:"end_time"`
	CheckinTime   *time.Time             `json:"check_in_time"`
	Room          string                 `json:"room"`
	CheckinStatus constant.CheckInStatus `json:"check_in_status"`
}

type CheckInHistoryListElement struct {
	ScheduleID    uint                   `json:"schedule_id"`
	Course        string                 `json:"course"`
	StartTime     time.Time              `json:"start_time"`
	EndTime       time.Time              `json:"end_time"`
	CheckinTime   *time.Time             `json:"check_in_time"`
	Room          string                 `json:"room"`
	CheckinStatus constant.CheckInStatus `json:"check_in_status"`
}

type ChangeAttendanceStatusRequest struct {
	ScheduleID           uint                          `json:"schedule_id" binding:"required"`
	RequestCheckInStatus constant.RequestCheckInStatus `json:"request_checkin_status" binding:"required"`
	ToUserID             uint                          `json:"to_user_id" binding:"required"`
	Reason               string                        `json:"reason" binding:"required"`
}

type AttendanceStatusInFormRequest struct {
	ID            uint                   `json:"id"`
	CheckInTime   *time.Time             `json:"checkin_time"`
	CheckInStatus constant.CheckInStatus `json:"checkin_status"`
}
