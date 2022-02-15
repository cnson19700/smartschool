package entity

import "gorm.io/gorm"

type Room struct {
	ID       int `gorm:"primaryKey autoCreateTime"`
	RoomID   string
	DeviceID string
	Name     string
	Courses  []*Course      `gorm:"many2many:schedulers; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsDelete gorm.DeletedAt `gorm:"column:deleted_at"`
}
