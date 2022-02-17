package entity

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	ID          int `gorm:"primaryKey autoCreateTime" json:"id"`
	StudentID   int `json:"student_id"`
	SchedulerID int `json:"scheduler_id"`
	// CourseID      int            `json:"course_id"`
	// RoomID        int            `json:"room_id"`
	// EndTime       time.Time      `json:"end_time"`
	CheckInTime   time.Time  `json:"checkin_time"`
	CheckInStatus string     `json:"checkin_status"`
	Student       *Student   `gorm:"foreignKey:ID;references:StudentID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Scheduler     *Scheduler `gorm:"foreignKey:ID;references:SchedulerID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Course        Course         `gorm:"foreignKey:ID;references:CourseID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Room          Room           `gorm:"foreignKey:ID;references:RoomID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
