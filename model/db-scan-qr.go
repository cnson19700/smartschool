package model

import (
	"time"

	"github.com/smartschool/model/entity"
	"gorm.io/gorm"
)

type EventScanCard struct {
	gorm.Model
	CardID    string
	Device    entity.Device
	DeviceID  uint `gorm:"index:device_id"`
	Company   entity.Company
	CompanyID uint `gorm:"index:company_id"`

	//Update field
	DepartmentID uint
	Department   Department
	Datetime     time.Time `gorm:"column:datetime"`
}
type ScanQRCodeData struct {
	CourseID        string `json:"courseId"`
	StudentID       string `json:"studentId"`
	Timestamp       int64  `json:"timeStamp"`
	CompanyTokenKey string `json:"companyTokenKey"`
}

func (mapping ScanQRCodeData) ConvertObjectToEventScanData() EventScanCard {
	device := Device{}.FindDeviceByToken(mapping.CompanyTokenKey)

	data := ScanQRCodeData{
		CardID:   mapping.CardID,
		DeviceID: device.ID,
		Datetime: time.Unix(ConvertDeviceTimestampToExact(mapping.Timestamp), 0),
	}

	return data
}
