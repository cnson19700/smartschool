package entity

import "gorm.io/gorm"

type Student struct {
	ID        uint   `gorm:"primaryKey; column:id" json:"id"`
	StudentID string `gorm:"index; column:student_id" json:"student_id"`
	Batch     string `gorm:"column:batch" json:"batch"`
	gorm.Model

	//User    *User     `gorm:"foreignKey:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Courses []*Course `gorm:"many2many:student_course_enrollments; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
