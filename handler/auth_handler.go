package handler

import (
	"net/http"

	"github.com/Nziza21/user-service/internal/service"
	notificationservice "github.com/Nziza21/user-service/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService  *service.AuthService
	emailService *notificationservice.SMTPEmailService
}

func NewAuthHandler(
	authService *service.AuthService,
	emailService *notificationservice.SMTPEmailService,
) *AuthHandler {
	return &AuthHandler{
		authService:  authService,
		emailService: emailService,
	}
}

// RequestResetPassword godoc
// @Summary      Request a password reset OTP
// @Description  Sends a one-time OTP to the user's email for password reset
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body body RequestResetPasswordInput true "Email of the user"
// @Success      200 {object} map[string]string "OTP sent to email"
// @Failure      400 {object} map[string]string "Bad request"
// @Failure      500 {object} map[string]string "Internal server error"
// @Failure      429 {object} map[string]string "Too many requests"
// @Router       /api/v1/auth/request-reset-password [post]
func (h *AuthHandler) RequestResetPassword(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request payload",
		})
		return
	}

	// Check OTP request rate limit
	allowed, err := h.authService.CheckOTPRequestLimit(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "rate limit check failed",
		})
		return
	}

	if !allowed {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error": "too many OTP requests, try again later",
		})
		return
	}

	// Check if user exists
	user, err := h.authService.UserRepo.FindByEmail(req.Email)
	if err != nil || user == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "OTP has been sent to your Email",
		})
		return
	}

	// Generate and save OTP
	otp := h.authService.GenerateOTP()
	if err := h.authService.SaveOTP(req.Email, otp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save OTP",
		})
		return
	}

	// Send OTP via email
	if err := h.emailService.SendOTPEmail(req.Email, otp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to send OTP email",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OTP has been sent to your Email",
	})
}

// ResetPassword godoc
// @Summary      Reset password
// @Description  Validate OTP and update the user's password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body body ResetPasswordInput true "Email, OTP, New Password"
// @Success      200 {object} map[string]string "Password reset successful"
// @Failure      400 {object} map[string]string "Bad request"
// @Failure      401 {object} map[string]string "Unauthorized - invalid or expired OTP"
// @Failure      404 {object} map[string]string "User not found"
// @Failure      500 {object} map[string]string "Internal server error"
// @Router       /api/v1/auth/reset-password [post]
func (h *AuthHandler) ResetPassword(c *gin.Context) {
    var req ResetPasswordInput
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if !h.authService.ValidateOTP(req.Email, req.OTP) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired OTP"})
        return
    }

    if err := h.authService.ResetPassword(req.Email, req.NewPassword); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "password reset failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "password reset successful"})
}