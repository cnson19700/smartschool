package dto

type UpdatePasswordRequest struct {
	Password  string `json:"password"`
	NewPass   string `json:"new_password"`
	ReNewPass string `json:"re_new_password"`
}
