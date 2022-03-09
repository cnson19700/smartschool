package entity

import "gorm.io/gorm"

type Student struct {
	//ID        uint           `gorm:"primaryKey autoIncrement column:id" json:"id"`
	ID        uint           `gorm:"primaryKey column:id" json:"id"`
	StudentID string         `gorm:"column:student_id" json:"student_id"`
	Batch     string         `gorm:"column:batch" json:"batch"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	User    *User     `gorm:"foreignKey:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Courses []*Course `gorm:"many2many:student_course_enrollments; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
