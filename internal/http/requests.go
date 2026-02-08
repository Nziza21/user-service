package http

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