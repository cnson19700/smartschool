package dto

type SemesterListElement struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Year  string `json:"year"`
}
