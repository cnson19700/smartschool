package entity

import "gorm.io/gorm"

type Course struct {
	ID       int `gorm:"primaryKey autoCreateTime"`
	CourseID string
	Name     string
	Students []*Student     `gorm:"many2many:student_courses; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Rooms    []*Room        `gorm:"many2many:schedulers; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsDelete gorm.DeletedAt `gorm:"column:deleted_at"`
}
