package entity

type User struct {
	ID       int    `gorm:"primaryKey autoCreateTime" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}