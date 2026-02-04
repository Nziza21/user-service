package http

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *UserHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{

		api.POST("/login", userHandler.Login)

		api.POST("/users", userHandler.CreateUser)
		api.GET("/users", userHandler.ListUsers)
		api.GET("/users/:id", userHandler.GetUserByID)
		api.PATCH("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)
	}

	return r
}