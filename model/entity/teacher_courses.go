package entity

import (
	"gorm.io/gorm"
)

type TeacherCourse struct {
	ID          uint   `gorm:"primaryKey; autoIncrement; column:id" json:"id"`
	TeacherID   uint   `gorm:"column:teacher_id"  json:"teacher_id"` //index of record
	CourseID    uint   `gorm:"column:course_id"  json:"course_id"`   //index of record
	TeacherRole string `gorm:"column:teacher_role" json:"teacher_role"`
	gorm.Model

	Teacher *Teacher `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course  *Course  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
