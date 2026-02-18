package handler

type CreateUserRequest struct {
	FullName string `json:"fullName" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

type UpdateProfileRequest struct {
	FullName string `json:"fullName,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type RequestResetPasswordInput struct {
    Email string `json:"email" binding:"required,email"`
}

type ResetPasswordInput struct {
    Email       string `json:"email" binding:"required,email"`
    OTP         string `json:"otp" binding:"required"`
    NewPassword string `json:"new_password" binding:"required"`
}