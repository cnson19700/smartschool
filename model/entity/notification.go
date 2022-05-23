package entity

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	// *BaseAppModel include Created, Updated fields and default callback methods
	gorm.Model
	ID           uint `gorm:"primary_key"`
	CreatedAt    time.Time
	Language     string `gorm:"type:varchar(3);index"`
	Os           string `gorm:"type:varchar(10);index"`
	Title        string
	Message      string
	Body         string
	Action       string
	ActionTitle  string
	Payload      string
	DirectAction bool
	Sent         bool
	Enabled      bool
	FilterObject string `gorm:"type:jsonb;default:'{\"all\":true}'"`
	Users        []User `gorm:"many2many:user_notification_mappings"`
}

func NewNotification() *Notification {
	return &Notification{}
}

// PUBLIC METHODS

func (p *Notification) PrimaryKey() []string {
	return []string{"NotificationID"}
}
