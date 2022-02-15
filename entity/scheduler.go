package entity

import (
	"time"

	"gorm.io/gorm"
)

type Scheduler struct {
	ID        int            `gorm:"primaryKey autoCreateTime" json:"id"`
	RoomID    int            `json:"room_id"`
	CourseID  int            `json:"course_id"`
	StartTime time.Time      `json:"start_time"`
	EndTime   time.Time      `json:"end_time"`
	Room      Room           `gorm:"foreignKey:RoomID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course    Course         `gorm:"foreignKey:CourseID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsDelete  gorm.DeletedAt `gorm:"column:deleted_at"`
}
