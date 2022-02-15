package entity

import (
	"time"

	"gorm.io/gorm"
)

type Scheduler struct {
	ID        int `gorm:"primaryKey autoCreateTime"`
	RoomID    int
	CourseID  int
	StartTime time.Time
	EndTime   time.Time
	Room      Room           `gorm:"foreignKey:RoomID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course    Course         `gorm:"foreignKey:CourseID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsDelete  gorm.DeletedAt `gorm:"column:deleted_at"`
}
