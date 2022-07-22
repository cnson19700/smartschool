package dto

import (
	"time"
)

type RequestChangeAttendanceForm struct {
	ScheduleID    uint       `json:"schedule_id"`
	CourseName    string     `json:"course_name"`
	StartTime     time.Time  `json:"start_time"`
	EndTime       time.Time  `json:"end_time"`
	Room          string     `json:"room"`
	CurrentStatus string     `json:"current_status"`
	CheckInTime   *time.Time `json:"checkin_time"`
}

type TeacherList struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type MobileViewComplainForm struct {
	FormID        uint      `json:"form_id"`
	CreatedTime   time.Time `json:"created_time"`
	CourseName    string    `json:"course_name"`
	ToTeacherName string    `json:"teacher_name"`
	CurrentStatus string    `json:"current_status"`
	RequestStatus string    `json:"request_status"`
	FormStatus    string    `json:"form_status"`
}

type MobileViewDetailComplainForm struct {
	CourseName    string     `json:"course_name"`
	StartTime     time.Time  `json:"start_time"`
	EndTime       time.Time  `json:"end_time"`
	Room          string     `json:"room"`
	CheckInTime   *time.Time `json:"checkin_time"`
	CurrentStatus string     `json:"current_status"`
	RequestStatus string     `json:"request_status"`
	ToTeacherName string     `json:"teacher_name"`
	FormStatus    string     `json:"form_status"`
	Reason        string     `json:"reason"`
	RejectReason  string     `json:"reject_reason"`
}
