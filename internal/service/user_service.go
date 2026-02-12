package service

import (
    "github.com/Nziza21/user-service/internal/Entities"
    "github.com/Nziza21/user-service/internal/store/repository"
    "golang.org/x/crypto/bcrypt"
    "github.com/google/uuid"
)

type UpdateUserReq struct {
    FullName string `json:"fullName"`
    Phone    string `json:"phone"`
}

type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *domain.User, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.PasswordHash = string(hashedPassword)
    return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id uuid.UUID) (*domain.User, error) {
    return s.repo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *domain.User) error {
    return s.repo.UpdateUser(user)
}

func (s *UserService) UpdateUserByID(id uuid.UUID, req UpdateUserReq) (*domain.User, error) {
    user, err := s.repo.GetUserByID(id)
    if err != nil {
        return nil, err
    }

    if req.FullName != "" {
        user.FullName = req.FullName
    }
    if req.Phone != "" {
        user.Phone = req.Phone
    }

    if err := s.repo.UpdateUser(user); err != nil {
        return nil, err
    }

    return user, nil
}

func (s *UserService) DeleteUser(user *domain.User) error {
    return s.repo.DeleteUser(user)
}

func (s *UserService) ListUsers(opts repository.ListUsersOpts) ([]domain.User, error) {
    return s.repo.ListUsers(opts) // user filters
}

func (s *UserService) DeleteUserByID(id uuid.UUID) error {
    user, err := s.repo.GetUserByID(id)
    if err != nil {
        return err
    }

    return s.repo.DeleteUser(user)
}

func (s *UserService) GetUserByEmail(email string) (*domain.User, error) {
    return s.repo.GetByEmail(email) 
}

func (s *UserService) CheckPassword(user *domain.User, password string) bool {
    return bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) == nil
}