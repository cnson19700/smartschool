package entity

import (
	"time"

	"gorm.io/gorm"
)

type DeviceSignalLog struct {
	CardId          string    `gorm:"column:card_id" json:"cardId"`
	Timestamp       time.Time `gorm:"column:timestamp" json:"timeStamp"`
	CompanyTokenKey string    `gorm:"column:company_token_key" json:"companyTokenKey"`
	Status          string    `gorm:"column:status" json:"status"`
	gorm.Model
}
