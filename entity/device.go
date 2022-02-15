package entity

import "gorm.io/gorm"

type Device struct {
	ID        int            `gorm:"primaryKey autoCreateTime" json:"id"`
	DeviceID  string         `json:"device_id"`
	RoomID    string         `json:"room_id"`
	Room      Room           `gorm:"foreignKey:RoomID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
