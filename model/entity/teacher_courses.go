package entity

import (
	"gorm.io/gorm"
)

type TeacherCourse struct {
	TeacherID uint `gorm:"primaryKey;column:teacher_id"  json:"teacher_id"` //index of record
	CourseID  uint `gorm:"primaryKey;column:course_id"  json:"course_id"`   //index of record
	gorm.Model

	Teacher *Teacher `gorm:"foreignKey:ID;references:TeacherID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course  *Course  `gorm:"foreignKey:ID;references:CourseID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
