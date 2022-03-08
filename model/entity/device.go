package entity

import "gorm.io/gorm"

type Device struct {
	ID        int            `gorm:"primaryKey autoCreateTime" json:"id"`
	DeviceID  string         `json:"device_id"`
	RoomID    int            `json:"room_id"`
	Room      Room           `gorm:"foreignKey:ID;references:RoomID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
