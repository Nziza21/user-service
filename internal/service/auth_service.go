package service

import (
    "fmt"
    "math/rand"
    "time"
    "errors"
    "github.com/redis/go-redis/v9"
    "golang.org/x/crypto/bcrypt"
    "strconv"
    "github.com/Nziza21/user-service/internal/store/cache"
    "github.com/Nziza21/user-service/internal/store/repository"
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
    r := rand.New(rand.NewSource(time.Now().UnixNano())) // Generates different tokens on different calls
    return fmt.Sprintf("%06d", r.Intn(1000000))
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

func (a *AuthService) CheckOTPRequestLimit(email string) (bool, error) {
    key := fmt.Sprintf("otp_req_limit:%s", email)

    countStr, err := a.RedisClient.Get(key)
    if err != nil {
        if errors.Is(err, redis.Nil) {
            // First request
            if err := a.RedisClient.Set(key, "1", 10*time.Minute); err != nil {
                return false, err
            }
            return true, nil
        }
        return false, err
    }

    count, _ := strconv.Atoi(countStr)

    if count >= 3 {
        return false, nil
    }

    count++
    if err := a.RedisClient.Set(key, strconv.Itoa(count), 10*time.Minute); err != nil {
        return false, err
    }

    return true, nil
}