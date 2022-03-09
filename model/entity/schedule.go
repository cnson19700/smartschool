package entity

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	ID        uint           `gorm:"primaryKey autoIncrement column:id" json:"id"`
	RoomID    uint           `gorm:"column:room_id" json:"room_id"`
	CourseID  uint           `gorm:"column:course_id" json:"course_id"`
	StartTime time.Time      `gorm:"column:start_time" json:"start_time"`
	EndTime   time.Time      `gorm:"column:end_time" json:"end_time"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	Room   *Room   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course *Course `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
