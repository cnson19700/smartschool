package entity

import "gorm.io/gorm"

type Device struct {
	ID        uint           `gorm:"primaryKey autoIncrement column:id" json:"id"`
	DeviceID  string         `gorm:"column:device_id" json:"device_id"`
	RoomID    uint           `gorm:"column:room_id" json:"room_id"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	Room *Room `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
