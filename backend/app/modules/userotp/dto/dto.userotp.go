package userotpdto

type CreateUserOtp struct {
	UserID string `json:"user_id"`
	Otp    string `json:"otp"`
}

type GetByIDUserOtp struct {
	ID string `uri:"id" binding:"required"`
}

type UserResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Otp       string `json:"otp"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
