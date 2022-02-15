package entity

import "gorm.io/gorm"

type Student struct {
	ID          int `gorm:"primaryKey autoCreateTime"`
	StudentID   string
	Name        string
	Email       string
	PhoneNumber string
	Courses     []*Course `gorm:"many2many:student_courses;"`
	IsDelete    gorm.DeletedAt  `gorm:"column:deleted_at"`
}
