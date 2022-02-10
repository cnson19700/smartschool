package api_model

import (
	"fmt"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%v", time.Time(t).Unix())
	return []byte(stamp), nil
}

type AuthenticationRes struct {
	Id    uint   `json:"id"`
	Phone string `json:"phone"`
}

type NotificationRes struct {
	Id        string    `json:"id"`
	CreatedAt *JSONTime `json:"createdAt"`
	Avatar    string    `json:"avatar"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Type      string    `json:"type"`
	Topic     string    `json:"topic"`
	Payload   string    `json:"payload"`
}
