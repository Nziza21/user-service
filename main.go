package main

import (
	"log"
	"strings"
    
	"github.com/Nziza21/user-service/internal/store/cache"
	"github.com/Nziza21/user-service/internal/store/config"
	"github.com/Nziza21/user-service/internal/db"
	myhttp "github.com/Nziza21/user-service/handler"
	"github.com/Nziza21/user-service/internal/store/repository"
	"github.com/Nziza21/user-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/Nziza21/user-service/docs"
)

var jwtSecret = []byte("mysecretpassword")  // Key used in Authentication

func AdminOnlyGin() gin.HandlerFunc { // Authentication Handler that passes to handler
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {  // If the header is empty
            c.JSON(401, gin.H{"error": "missing token"}) // Error if authentication fails
            c.Abort()
            return
        }

        tokenStr := strings.TrimPrefix(authHeader, "Bearer ") // string token
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        }) // Verifies token with jwt secret key(mysecretkey), returns the token

        if err != nil || !token.Valid {
            c.JSON(401, gin.H{"error": "invalid token"}) // If parsing fails, the process stops and get an error message
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims) // Claims : the request
        if !ok || claims["role"] != "admin" { // If role is not admin
            c.JSON(403, gin.H{"error": "admin access required"}) // returns this if the role isn't admin. 403 FORBIDEN
            c.Abort()
            return
        }

        c.Next() // If everything checks out, the request goes to the handler
    }
}

func main() {
    cfg := config.LoadConfig() // configuration for all the settings of the App

    // Redis configuration
    redisClient := cache.NewRedisClient("localhost:6379", "", 0)

    // DB connection 
    database := db.ConnectDB(cfg.DB_DSN)
    log.Println("Database connected:", database != nil) 

    // Repos & services
    userRepo := repository.NewUserRepository(database) // repo that connects to the DB
    userService := service.NewUserService(userRepo) // business logic layer that returns repo for CRUD ops
    userHandler := myhttp.NewUserHandler(userService) // Sends request via GIN -> business logic
    authService := service.NewAuthService(userRepo, redisClient) // To retrieve DB info & stores OTP
    emailService := service.NewSMTPEmailService() // email tool fpr sending OPTs
    authHandler := myhttp.NewAuthHandler(authService, emailService) // For request feedback

    // Gin router
    r := gin.Default()
    r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    v1 := r.Group("/api/v1/users") // grouping routes path
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