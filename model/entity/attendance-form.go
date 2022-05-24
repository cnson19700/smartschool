package entity

import (
	"github.com/smartschool/lib/constant"
	"gorm.io/gorm"
)

type AttendanceForm struct {
	ID                   uint                          `gorm:"primaryKey; autoIncrement; column:id" json:"id"`
	RequestUserID        uint                          `gorm:"column:request_user_id" json:"request_user_id"`
	ReceiveUserID        uint                          `gorm:"column:receive_user_id" json:"receive_user_id"`
	SemesterID           uint                          `gorm:"column:semester_id" json:"semester_id"`
	RequestStatus        constant.RequestStatus        `gorm:"column:request_status; type:request_status; default:'Pending'" json:"request_status"`
	RequestCheckInStatus constant.RequestCheckInStatus `gorm:"column:request_checkin_status; type:request_checkin_status; default:'Attend'" json:"request_checkin_status"`
	Reason               string                        `gorm:"column:reason" json:"reason"`
	RejectReason         string                        `gorm:"column:reject_reason" json:"reject_reason"`
	//	ScheduleID           uint                          `gorm:"column:schedule_id" json:"schedule_id"`
	//	AttendanceID         uint                          `gorm:"column:attendance_id" json:"attendance_id"`
	//	CheckInStatus        constant.CheckInStatus        `gorm:"column:checkin_status; type:checkin_status; default:'Unknown'" json:"checkin_status"`

	gorm.Model

	RequestUser *User        `gorm:"foreignKey:RequestUserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReceiveUser *User        `gorm:"foreignKey:ReceiveUserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Semester    *Semester    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Schedules   []Schedule   `gorm:"many2many:form_details; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Details     []FormDetail `gorm:"foreignKey:FormID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Schedule    *Schedule    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Attendance  *Attendance  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
