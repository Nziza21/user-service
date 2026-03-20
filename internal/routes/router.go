package routes

import (

	"github.com/Nziza21/user-service/handler"
	"github.com/gin-gonic/gin"

)


func SetupRouter(
    userHandler *handler.UserHandler,
    authHandler *handler.AuthHandler,
    productHandler *handler.ProductHandler,
    jwtSecret []byte,
) *gin.Engine {
    r := gin.Default()
    api := r.Group("/api/v1")

    setupAuthRoutes(api, userHandler, authHandler)
    setupUserRoutes(api, userHandler, jwtSecret)
    setupProductRoutes(api, productHandler) 

    return r
}