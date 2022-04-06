package dto

type CourseReportListElement struct {
	ID         uint   `json:"id"`
	CourseID   string `json:"course_id"`
	Name       string `json:"name"`
	Attendance uint   `json:"attendance"`
	Absence    uint   `json:"absence"`
	Total      uint   `json:"total"`
}

type CourseReportPartElement struct {
	ID       uint   `json:"id"`
	CourseID string `json:"course_id"`
	Name     string `json:"name"`
}
