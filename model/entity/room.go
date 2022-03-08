package entity

import "gorm.io/gorm"

type Room struct {
	ID        uint           `gorm:"primaryKey autoIncrement column:id" json:"id"`
	RoomID    string         `gorm:"column:room_id" json:"room_id"`
	Name      string         `gorm:"column:name" json:"name"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	Courses []*Course `gorm:"many2many:schedules; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
