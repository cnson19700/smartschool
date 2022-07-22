package entity

import (
	"gorm.io/gorm"
)

type ComplainForm struct {
	ID            uint   `gorm:"primaryKey; autoIncrement; column:id" json:"id"`
	RequestUserID uint   `gorm:"column:request_user_id" json:"request_user_id"`
	ReceiveUserID uint   `gorm:"column:receive_user_id" json:"receive_user_id"`
	SemesterID    uint   `gorm:"column:semester_id" json:"semester_id"`
	RequestStatus string `gorm:"column:request_status" json:"request_status"`
	FormStatus    string `gorm:"column:form_status" json:"form_status"`
	Reason        string `gorm:"column:reason" json:"reason"`
	RejectReason  string `gorm:"column:reject_reason" json:"reject_reason"`
	ScheduleID    uint   `gorm:"column:schedule_id" json:"schedule_id"`

	gorm.Model

	RequestUser *User     `gorm:"foreignKey:RequestUserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReceiveUser *User     `gorm:"foreignKey:ReceiveUserID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Semester    *Semester `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Schedule    *Schedule `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
