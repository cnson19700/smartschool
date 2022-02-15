package entity

import "time"

type DeviceSignal struct {
	CardId          string    `json:"cardId"`
	TimeStamp       time.Time `json:"timeStamp"`
	CompanyTokenKey string    `json:"companyTokenKey"`
}
