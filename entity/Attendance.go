package entity

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	ID            int `gorm:"primaryKey autoCreateTime"`
	StudentID     string
	CourseID      string
	RoomID        string
	CheckInTime   time.Time
	EndTime       time.Time
	CheckInStatus string
	Student       Student        `gorm:"foreignKey:StudentID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course        Course         `gorm:"foreignKey:CourseID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Room          Room           `gorm:"foreignKey:RoomID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsDelete      gorm.DeletedAt `gorm:"column:deleted_at"`
}
