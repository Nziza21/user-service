package http

import (
    "net/http"

    "github.com/Nziza21/user-service/internal/domain"
    "github.com/Nziza21/user-service/internal/service"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/golang-jwt/jwt/v4"
)

type UserHandler struct {
    userService *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
    return &UserHandler{userService: s}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user in the system
// @Tags Users
// @Accept json
// @Produce json
// @Param user body domain.User true "User Data"
// @Success 201 {object} domain.User
// @Failure 400 {object} http.ErrorResponse
// @Failure 404 {object} http.ErrorResponse
// @Failure 500 {object} http.ErrorResponse
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
    var req struct {
        FullName string `json:"fullName" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Phone    string `json:"phone"`
        Password string `json:"password" binding:"required"`
        Role     string `json:"role"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := &domain.User{
        FullName: req.FullName,
        Email:    req.Email,
        Phone:    req.Phone,
        Role:     req.Role,
    }

    if err := h.userService.CreateUser(user, req.Password); err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusCreated, user)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get user details by user ID
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} domain.User
// @Failure 400 {object} http.ErrorResponse
// @Failure 404 {object} http.ErrorResponse
// @Failure 500 {object} http.ErrorResponse
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := uuid.Parse(idParam)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

    user, err := h.userService.GetUserByID(id)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    c.IndentedJSON(http.StatusOK, user)
}

// ListUsers godoc
// @Summary List all users
// @Description Retrieve a list of all users (admin only)
// @Tags Users
// @Produce json
// @Success 200 {array} domain.User
// @Failure 400 {object} http.ErrorResponse
// @Failure 404 {object} http.ErrorResponse
// @Failure 500 {object} http.ErrorResponse
// @Router /api/v1/users [get]
// @Security ApiKeyAuth
func (h *UserHandler) ListUsers(c *gin.Context) {
    users, err := h.userService.ListUsers()
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update user details by user ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body service.UpdateUserReq true "Updated User Data"
// @Success 200 {object} domain.User
// @Failure 400 {object} http.ErrorResponse
// @Failure 404 {object} http.ErrorResponse
// @Failure 500 {object} http.ErrorResponse
// @Router /api/v1/users/{id} [patch]
func (h *UserHandler) UpdateUser(c *gin.Context) {
    idParam := c.Param("id")
    id, err := uuid.Parse(idParam)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

    var req service.UpdateUserReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedUser, err := h.userService.UpdateUserByID(id, req)
    if err != nil {
        if err.Error() == "user not found" {
            c.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
        } else {
            c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.IndentedJSON(http.StatusOK, updatedUser)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by user ID (admin only)
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} http.MessageResponse
// @Router /api/v1/users/{id} [delete]
// @Security ApiKeyAuth
func (h *UserHandler) DeleteUser(c *gin.Context) {
    idParam := c.Param("id")
    id, err := uuid.Parse(idParam)
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

    if err := h.userService.DeleteUserByID(id); err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

// LoginUser godoc
// @Summary User login
// @Description Login with email and password, returns JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body struct{Email string; Password string} true "Credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} http.ErrorResponse
// @Failure 401 {object} http.ErrorResponse
// @Router /api/v1/auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {

	var req struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetUserByEmail(req.Email)
	if err != nil || !h.userService.CheckPassword(user, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"role":    user.Role,
	})
	tokenString, _ := token.SignedString(jwtSecret)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}