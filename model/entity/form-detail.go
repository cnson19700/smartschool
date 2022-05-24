package entity

import (
	"gorm.io/gorm"
)

type FormDetail struct {
	// ID         uint `gorm:"primaryKey; autoIncrement; column:id" json:"id"`
	AttendanceFormID uint `gorm:"column:attendance_form_id" json:"attendance_form_id"`
	ScheduleID       uint `gorm:"column:schedule_id" json:"schedule_id"`
	//	AttendanceID uint `gorm:"column:attendance_id" json:"attendance_id"`
	gorm.Model

	// Schedule   *Schedule                          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Form       *ChangeAttendanceStatusFormRequest `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Attendance *Attendance                        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
