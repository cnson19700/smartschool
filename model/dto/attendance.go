package dto

import "time"

type AttendanceListElement struct {
	StartTime     time.Time  `json:"start_time"`
	EndTime       time.Time  `json:"end_time"`
	CheckinTime   *time.Time `json:"check_in_time"`
	Room          string     `json:"room"`
	CheckinStatus string     `json:"check_in_status"`
	ScheduleID    uint       `json:"schedule_id"`
}

type CheckInHistoryListElement struct {
	Course        string     `json:"course"`
	StartTime     time.Time  `json:"start_time"`
	EndTime       time.Time  `json:"end_time"`
	CheckinTime   *time.Time `json:"check_in_time"`
	Room          string     `json:"room"`
	CheckinStatus string     `json:"check_in_status"`
	ScheduleID    uint       `json:"schedule_id"`
}

type ChangeAttendanceStatusRequest struct {
	ScheduleID           uint   `json:"schedule_id" binding:"required"`
	RequestCheckInStatus string `json:"request_checkin_status" binding:"required"`
	ToUserID             uint   `json:"to_user_id" binding:"required"`
	Reason               string `json:"reason" binding:"required"`
}

type AttendanceStatusInFormRequest struct {
	ID            uint       `json:"id"`
	CheckInTime   *time.Time `json:"checkin_time"`
	CheckInStatus string     `json:"checkin_status"`
}
