package router

import (
	"github.com/gin-gonic/gin"
	 "github.com/Nziza21/user-service/handler"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
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