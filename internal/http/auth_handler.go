package http

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/Nziza21/user-service/internal/service"
	notificationservice "github.com/Nziza21/user-service/notification-service"
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

func (h *AuthHandler) RequestResetPassword(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	otp := h.authService.GenerateOTP()

	if err := h.authService.SaveOTP(req.Email, otp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save OTP"})
		return
	}

	if err := h.emailService.SendOTPEmail(req.Email, otp); err != nil {
	log.Println("email error:", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send OTP"})
	return
    }


	c.JSON(http.StatusOK, gin.H{"message": "OTP sent to email"})
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req struct {
		Email       string `json:"email" binding:"required,email"`
		OTP         string `json:"otp" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !h.authService.ValidateOTP(req.Email, req.OTP) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired OTP"})
		return
	}

	user, err := h.authService.UserRepo.GetByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password hashing failed"})
		return
	}

	user.PasswordHash = string(hash)

	if err := h.authService.UserRepo.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password reset successful"})
}