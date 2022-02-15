package dto

type EventCheckin struct {
	StudentId string `json:"studentId"`
	Timestamp int64  `json:"timestamp"`
	Location  string `json:"location"`
}
