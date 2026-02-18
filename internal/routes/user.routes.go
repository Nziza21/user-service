package routes

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/Nziza21/user-service/handler"
	"github.com/golang-jwt/jwt/v4"
)

func AdminOnlyGin(jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["role"] != "admin" {
			c.JSON(403, gin.H{"error": "admin access required"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func setupUserRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler, jwtSecret []byte) {
	users := api.Group("/users")
	users.POST("", userHandler.CreateUser)
	users.GET("", AdminOnlyGin(jwtSecret), userHandler.ListUsers)
	users.GET("/:id", userHandler.GetUserByID)
	users.PATCH("/:id", userHandler.UpdateUser)
	users.DELETE("/:id", AdminOnlyGin(jwtSecret), userHandler.DeleteUser)
}