package entity

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	ID            uint      `gorm:"primaryKey; autoIncrement; column:id" json:"id"`
	UserID        uint      `gorm:"index; column:user_id" json:"user_id"`
	ScheduleID    uint      `gorm:"index; column:schedule_id" json:"schedule_id"`
	CheckInTime   time.Time `gorm:"column:checkin_time" json:"checkin_time"`
	CheckInStatus string    `gorm:"column:checkin_status" json:"checkin_status"`
	//DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`
	gorm.Model

	User     *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Schedule *Schedule `gorm:"foreignKey:ID;references:ScheduleID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
