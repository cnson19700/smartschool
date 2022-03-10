package entity

import "gorm.io/gorm"

type Role struct {
	ID    uint   `gorm:"primaryKey; autoIncrement; column:id" json:"id"`
	Title string `gorm:"column:title" json:"title"`
	//DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	gorm.Model

	//User []*User `gorm:"foreignKey:RoleID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
