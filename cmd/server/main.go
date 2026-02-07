package main

import (
	"log"
	"strings"

	"github.com/Nziza21/user-service/internal/cache"
	"github.com/Nziza21/user-service/internal/config"
	"github.com/Nziza21/user-service/internal/db"
	myhttp "github.com/Nziza21/user-service/internal/http"
	"github.com/Nziza21/user-service/internal/repository"
	"github.com/Nziza21/user-service/internal/service"
	notificationservice "github.com/Nziza21/user-service/notification-service"

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

    // Redis
    redisClient := cache.NewRedisClient("localhost:6379", "", 0)

    // DB
    database := db.ConnectDB(cfg.DB_DSN)
    log.Println("Database connected:", database != nil)

    // Repos & services
    userRepo := repository.NewUserRepository(database)
    userService := service.NewUserService(userRepo)
    userHandler := myhttp.NewUserHandler(userService)
    authService := service.NewAuthService(userRepo, redisClient)
    emailService := notificationservice.NewSMTPEmailService()
    authHandler := myhttp.NewAuthHandler(authService, emailService)



    // Gin router
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

    // Auth routes
    r.POST("/api/v1/auth/login", userHandler.Login)
    r.POST("/api/v1/auth/request-reset", authHandler.RequestResetPassword)
    r.POST("/api/v1/auth/reset-password", authHandler.ResetPassword)

    log.Println("Starting server on port", cfg.Port)
    if err := r.Run(":" + cfg.Port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}