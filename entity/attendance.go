package entity

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	ID            int            `gorm:"primaryKey autoCreateTime" json:"id"`
	StudentID     int            `json:"student_id"`
	CourseID      int            `json:"course_id"`
	RoomID        int            `json:"room_id"`
	CheckInTime   time.Time      `json:"checkin_time"`
	EndTime       time.Time      `json:"end_time"`
	CheckInStatus string         `json:"checkin_status"`
	Student       Student        `gorm:"foreignKey:ID;references:StudentID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course        Course         `gorm:"foreignKey:ID;references:CourseID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Room          Room           `gorm:"foreignKey:ID;references:RoomID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`
}
