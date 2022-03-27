package entity

import "gorm.io/gorm"

type Teacher struct {
	ID        uint   `gorm:"primaryKey; column:id" json:"id"`
	TeacherID string `gorm:"index; column:teacher_id" json:"teacher_id"`
	gorm.Model

	Courses []*Course `gorm:"many2many: teachers_courses; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
