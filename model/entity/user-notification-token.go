package entity

import "gorm.io/gorm"

type UserNotificationToken struct {
	ID     uint   `gorm:"primaryKey; autoIncrement; column:id" json:"id"`
	UserID uint   `gorm:"column:user_id" json:"user_id"`
	Token  string `gorm:"column:token" json:"token"`
	gorm.Model
}
