package entity

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	ID            int            `gorm:"primaryKey autoCreateTime" json:"id"`
	StudentID     string         `json:"student_id"`
	CourseID      string         `json:"course_id"`
	RoomID        string         `json:"room_id"`
	CheckInTime   time.Time      `json:"checkin_time"`
	EndTime       time.Time      `json:"end_time"`
	CheckInStatus string         `json:"checkin_status"`
	Student       Student        `gorm:"foreignKey:StudentID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course        Course         `gorm:"foreignKey:CourseID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Room          Room           `gorm:"foreignKey:RoomID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`
}
