package entity

import "gorm.io/gorm"

type Course struct {
	ID              uint   `gorm:"primaryKey; autoIncrement; column:id" json:"id"`
	CourseID        string `gorm:"column:course_id" json:"course_id"`
	TeacherID       string `gorm:"column:teacher_id" json:"teacher_id"`
	TeacherRole     string `gorm:"column:teacher_role" json:"teacher_role"`
	Name            string `gorm:"column:name" json:"name"`
	NumberOfStudent int    `gorm:"column:number_of_student" json:"number_of_student"`
	SemesterID      uint   `gorm:"column:semester_id" json:"semester_id"`
	//DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at"`
	gorm.Model

	Semester *Semester  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Students []*Student `gorm:"many2many:student_course_enrollments; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Teacher  []*Teacher `gorm:"many2many:teacher_courses; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Rooms    []*Room    `gorm:"many2many:schedules; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CourseByTeacher struct {
	ID              uint   `json:"id"`
	CourseID        string `json:"course_id"`
	TeacherID       string `json:"teacher_id"`
	TeacherRole     string `json:"teacher_role"`
	Name            string `json:"course_name"`
	SemesterID      uint   `json:"semester_id"`
	NumberOfStudent int    `json:"number_of_students"`
}
