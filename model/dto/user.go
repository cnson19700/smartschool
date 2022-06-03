package dto

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordFirstTimeRequest struct {
	NewPass   string `json:"new_password"`
	ReNewPass string `json:"re_new_password"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordEmailData struct {
	NewPassword string
}
