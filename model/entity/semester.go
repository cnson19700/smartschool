package entity

import (
	"time"

	"gorm.io/gorm"
)

type Semester struct {
	ID        uint      `gorm:"primaryKey; autoIncrement; column:id" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	Year      string    `gorm:"column:year" json:"year"`
	FacultyID uint      `gorm:"faculty_id" json:"faculty_id"`
	StartTime time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`
	//DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	gorm.Model

	Faculty *Faculty  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Courses []*Course `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
