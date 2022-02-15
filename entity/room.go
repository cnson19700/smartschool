package entity

import "gorm.io/gorm"

type Room struct {
	ID       int            `gorm:"primaryKey autoCreateTime" json:"id"`
	RoomID   string         `json:"room_id"`
	Name     string         `json:"name"`
	Courses  []*Course      `gorm:"many2many:schedulers; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsDelete gorm.DeletedAt `gorm:"column:deleted_at"`
}
