package dto

import "time"

type CheckInHistoryElement struct {
	CourseName    string    `json:"course_name"`
	CheckinTime   time.Time `json:"check_in_time"`
	CheckinStatus string    `json:"check_in_status"`
}
