package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint      `gorm:"primaryKey; autoIncrement; column:id" json:"id"`
	Username    string    `gorm:"column:user_name" json:"user_name"`
	Password    string    `gorm:"column:password" json:"password"`
	Email       string    `gorm:"column:email" json:"email"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number"`
	FirstName   string    `gorm:"firstname" json:"firstname"`
	LastName    string    `gorm:"lastname" json:"lastname"`
	DateOfBirth time.Time `gorm:"date_of_birth" json:"date_of_birth"`
	Gender      int       `gorm:"gender" json:"gender"`
	RoleID      uint      `gorm:"role_id" json:"role_id"`
	FacultyID   uint      `gorm:"faculty_id" json:"faculty_id"`
	IsActivate  bool      `gorm:"is_activate" json:"is_activate"`
	//DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
	gorm.Model

	Role    *Role    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Faculty *Faculty `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Student *Student `gorm:"foreignKey:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Teacher *Teacher `gorm:"foreignKey:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
