package service

import (
    "fmt"
    "math/rand"
    "time"
    "golang.org/x/crypto/bcrypt"

    "github.com/Nziza21/user-service/internal/cache"
    "github.com/Nziza21/user-service/internal/repository"
)

type AuthService struct {
    UserRepo    *repository.UserRepository
    RedisClient *cache.RedisClient
}

func NewAuthService(userRepo *repository.UserRepository, redisClient *cache.RedisClient) *AuthService {
    return &AuthService{
        UserRepo:    userRepo,
        RedisClient: redisClient,
    }
}

func (a *AuthService) GenerateOTP() string {
    rand.Seed(time.Now().UnixNano())
    return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func (a *AuthService) SaveOTP(userEmail, otp string) error {
    key := fmt.Sprintf("otp:%s", userEmail)
    return a.RedisClient.Set(key, otp, 5*time.Minute)
}

func (a *AuthService) ValidateOTP(userEmail, otp string) bool {
    key := fmt.Sprintf("otp:%s", userEmail)
    val, err := a.RedisClient.Get(key)
    if err != nil || val != otp {
        return false
    }
    _ = a.RedisClient.Delete(key)
    return true
}

func hashPassword(password string) string {
    hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hashed)
}