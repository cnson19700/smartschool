package entity

import "gorm.io/gorm"

type Faculty struct {
	ID        uint           `gorm:"primaryKey autoIncrement column:id" json:"id"`
	Title     string         `gorm:"column:title" json:"title"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	//Users     []*User     `gorm:"foreignKey:FacultyID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	//Semesters []*Semester `gorm:"foreignKey:FacultyID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
