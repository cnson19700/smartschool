package dto

type RegisterRequest struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Token    string `json:"token"`
	IP       string `json:"ip"`
}
