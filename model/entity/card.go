package entity

import "gorm.io/gorm"

type Card struct {
	ID     uint   `gorm:"primaryKey; autoIncrement; column:id" json:"id"`
	CardID string `gorm:"index; column:card_id" json:"card_id"`
	UserID uint   `gorm:"column:user_id" json:"user_id"`
	//DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	gorm.Model

	User *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
