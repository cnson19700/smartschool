package entity

import "gorm.io/gorm"

type ChangeAttendanceStatusRequest struct {
	ID               uint   `gorm:"primaryKey; autoIncrement; column:id" json:"id"`
	RequestUserID    uint   `gorm:"column:request_user_id" json:"request_user_id"`
	ReceiveUserID    uint   `gorm:"column:receive_user_id" json:"receive_user_id"`
	ScheduleID       uint   `gorm:"column:schedule_id" json:"schedule_id"`
	SemesterID       uint   `gorm:"column:semester_id" json:"semester_id"`
	Status           string `gorm:"column:status; type:enum('approve', 'reject', 'pending'); default:'pending'" json:"status"`
	OldCheckInStatus string `gorm:"column:old_checkin_status; type:enum('Late', 'Attend', 'Absence', 'Unknown'); default:'Unknown'" json:"old_checkin_status"`
	NewCheckInStatus string `gorm:"column:new_checkin_status; type:enum('Late', 'Attend', 'Absence with permission'); default:'Attend'" json:"new_checkin_status"`

	Reason       string `gorm:"column:reason" json:"reason"`
	RejectReason string `gorm:"column:reject_reason" json:"reject_reason"`

	gorm.Model

	RequestUser *User     `gorm:"foreignKey:RequestUserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReceiveUser *User     `gorm:"foreignKey:ReceiveUserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Schedule    *Schedule `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Semester    *Semester `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
