package entity

import "gorm.io/gorm"

type StudentCourseEnrollment struct {
	ID        uint           `gorm:"primaryKey autoIncrement column:id" json:"id"`
	StudentID uint           `gorm:"column:student_id" json:"student_id"`
	CourseID  uint           `gorm:"column:course_id" json:"course_id"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	Student *Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course  *Course  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
