// @title           User Service API
// @version         1.0
// @description     This is a user management service API.
// @termsOfService  http://example.com/terms/

// @contact.name   Nziza Samuel
// @contact.url    http://github.com/Nziza21
// @contact.email  nziza@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8800
// @BasePath 

package main

import (
    
    "log"
    
    "strings"

    "github.com/Nziza21/user-service/internal/config"
    "github.com/Nziza21/user-service/internal/db"
    myhttp "github.com/Nziza21/user-service/internal/http"
    "github.com/Nziza21/user-service/internal/repository"
    "github.com/Nziza21/user-service/internal/service"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
    _ "github.com/Nziza21/user-service/docs"
)

var jwtSecret = []byte("mysecretpassword") 

func AdminOnlyGin() gin.HandlerFunc {
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

func main() {
    cfg := config.LoadConfig()

    database := db.ConnectDB(cfg.DB_DSN)
    log.Println("Database connected:", database != nil)

    userRepo := repository.NewUserRepository(database)
    userService := service.NewUserService(userRepo)
    userHandler := myhttp.NewUserHandler(userService)

    r := gin.Default()

    r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    v1 := r.Group("/api/v1/users")
{
    v1.POST("", userHandler.CreateUser)
    v1.GET("", AdminOnlyGin(), userHandler.ListUsers)
    v1.GET("/:id", userHandler.GetUserByID)
    v1.PATCH("/:id", userHandler.UpdateUser)
    v1.DELETE("/:id", AdminOnlyGin(), userHandler.DeleteUser)
}
r.POST("/api/v1/auth/login", userHandler.Login)

    log.Println("Starting server on port", cfg.Port)
    if err := r.Run(":" + cfg.Port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}