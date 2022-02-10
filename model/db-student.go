package model

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	ID        string `json:"id"`
	Name      string `json:"name"`
	StudentID string `json:"student_id"`
}
