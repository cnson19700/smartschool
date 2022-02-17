package dto

import "time"

type HistoryElement struct {
	CheckinTime   time.Time `json:"check_in_time"`
	CheckinStatus string    `json:"check_in_status"`
}
