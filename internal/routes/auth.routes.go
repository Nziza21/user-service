package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Nziza21/user-service/handler"
)

func setupAuthRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler, authHandler *handler.AuthHandler) {
	auth := api.Group("/auth")
	auth.POST("/login", userHandler.Login)
	auth.POST("/request-reset-password", authHandler.RequestResetPassword)
	auth.POST("/reset-password", authHandler.ResetPassword)
}