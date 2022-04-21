package dto

import "time"

type AttendanceListElement struct {
	StartTime     time.Time  `json:"start_time"`
	EndTime       time.Time  `json:"end_time"`
	CheckinTime   *time.Time `json:"check_in_time"`
	Room          string     `json:"room"`
	CheckinStatus string     `json:"check_in_status"`
}

type CheckInHistoryListElement struct {
	Course        string     `json:"course"`
	StartTime     time.Time  `json:"start_time"`
	EndTime       time.Time  `json:"end_time"`
	CheckinTime   *time.Time `json:"check_in_time"`
	Room          string     `json:"room"`
	CheckinStatus string     `json:"check_in_status"`
}
