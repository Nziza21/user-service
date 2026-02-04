package repository

import (
    "github.com/Nziza21/user-service/internal/domain"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
    return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*domain.User, error) {
    var user domain.User
    result := r.db.First(&user, "id = ?", id)
    return &user, result.Error
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
    return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUser(user *domain.User) error {
    return r.db.Delete(user).Error
}

func (r *UserRepository) ListUsers() ([]domain.User, error) {
    var users []domain.User
    result := r.db.Find(&users)
    return users, result.Error
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
    var user domain.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}