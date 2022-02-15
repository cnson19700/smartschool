package entity

import "gorm.io/gorm"

type StudentCourse struct {
	ID        int `gorm:"primaryKey autoCreateTime"`
	StudentID int
	CourseID  int
	Student   Student        `gorm:"foreignKey:StudentID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course    Course         `gorm:"foreignKey:CourseID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
