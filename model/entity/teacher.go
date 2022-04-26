package entity

import "gorm.io/gorm"

type Teacher struct {
	ID        uint   `gorm:"primaryKey; column:id" json:"id"`
	TeacherID string `gorm:"index; column:teacher_id" json:"teacher_id"` //identified ID
	gorm.Model

	TeacherCourses []*TeacherCourse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
