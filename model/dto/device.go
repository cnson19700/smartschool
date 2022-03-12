package dto

type DeviceSignal struct {
	CardId          string `json:"cardId"`
	Timestamp       int64  `json:"timeStamp"`
	CompanyTokenKey string `json:"companyTokenKey"`
}

type EventCheckin struct {
	StudentId string `json:"studentId"`
	Timestamp int64  `json:"timestamp"`
	Location  string `json:"location"`
}
