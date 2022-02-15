package entity

import "gorm.io/gorm"

type Course struct {
	ID              int            `gorm:"primaryKey autoCreateTime" json:"id"`
	CourseID        string         `json:"course_id"`
	Name            string         `json:"name"`
	NumberOfStudent int            `json:"number_of_student"`
	Students        []*Student     `gorm:"many2many:student_courses; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Rooms           []*Room        `gorm:"many2many:schedulers; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsDelete        gorm.DeletedAt `gorm:"column:deleted_at"`
}
