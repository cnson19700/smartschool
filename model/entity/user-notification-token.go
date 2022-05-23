package entity

import (
	"gorm.io/gorm"
)

type UserNotificationToken struct {
	ID             uint   `gorm:"primaryKey; autoIncrement; column:id" json:"id"`
	UserID         uint   `gorm:"column:user_id" json:"user_id"`
	NotificationID uint   `gorm:"column:notification_id" json:"notification_id"`
	Token          string `gorm:"column:token" json:"token"`
	ReadStatus     bool   `gorm:"type:boolean;default:false"`
	User           User
	Notification   Notification
	gorm.Model
}
