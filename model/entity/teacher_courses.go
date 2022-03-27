package entity

type TeacherCourse struct {
	ID        int64 `gorm:"column:id;primary_key"  json:"id"`
	TeacherID int64 `gorm:"column:teacher_id"  json:"teacher_id"`
	CourseID  int64 `gorm:"column:course_id"  json:"course_id"`

	Teacher *Teacher `gorm:"foreignKey:ID;references:TeacherID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course  *Course  `gorm:"foreignKey:ID;references:CourseID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
