package entity

import "gorm.io/gorm"

type Faculty struct {
	ID    uint   `gorm:"primaryKey; autoIncrement; column:id" json:"id"`
	Title string `gorm:"column:title" json:"title"`
	//DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	gorm.Model

	Semesters []*Semester `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
