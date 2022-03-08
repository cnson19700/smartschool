package entity

import "gorm.io/gorm"

type Student struct {
	ID          int            `gorm:"primaryKey autoCreateTime" json:"id"`
	StudentID   string         `json:"student_id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	PhoneNumber string         `json:"phone_number"`
	Courses     []*Course      `gorm:"many2many:student_courses;"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}
