package dto

import "time"

type CheckInHistoryElement struct {
	CourseName    string    `json:"course_name"`
	CheckinTime   time.Time `json:"check_in_time"`
	CheckinStatus string    `json:"check_in_status"`
}

type StudentProfile struct {
	Name        string `json:"student_name"`
	Class       string `json:"student_class"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	StudentID   string `json:"student_id"`
}
