package authdto

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}
type VerifyOtpRequest struct {
	ID              string `json:"id"`
	Otp             string `json:"otp"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}
