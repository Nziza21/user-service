package main

import (
	"log"
	"net/http"

	_ "github.com/Nziza21/user-service/docs"
	myhttp "github.com/Nziza21/user-service/handler"
	"github.com/Nziza21/user-service/internal/db"
	router "github.com/Nziza21/user-service/internal/routes"
	"github.com/Nziza21/user-service/internal/service"
	"github.com/Nziza21/user-service/internal/store/cache"
	"github.com/Nziza21/user-service/internal/store/config"
	"github.com/Nziza21/user-service/internal/store/repository"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var jwtSecret = []byte("mysecretpassword")

//	@title						USER-SERVICE
//	@description				user-service
//	@version					1.2
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
func main() {
	cfg := config.LoadConfig()

	redisClient := cache.NewRedisClient("localhost:6379", "", 0)
	database := db.ConnectDB(cfg.DB_DSN)
	log.Println("Database connected:", database != nil)

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := myhttp.NewUserHandler(userService)
	authService := service.NewAuthService(userRepo, redisClient)
	emailService := service.NewSMTPEmailService()
	authHandler := myhttp.NewAuthHandler(authService, emailService)

	r := router.SetupRouter(userHandler, authHandler, jwtSecret)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/docs", func(c *gin.Context) { 
		c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})
	// API v1
	log.Println("Starting server on port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
