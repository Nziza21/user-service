package routes

import (

	"github.com/Nziza21/user-service/handler"
	"github.com/gin-gonic/gin"

)


func SetupRouter(userHandler *handler.UserHandler, authHandler *handler.AuthHandler, jwtSecret []byte) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")

	// Auth routes
	setupAuthRoutes(api, userHandler, authHandler)  

	// User routes
	setupUserRoutes(api, userHandler, jwtSecret)

	return r
}