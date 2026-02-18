package repository

import (
	"github.com/Nziza21/user-service/internal/Entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
    "errors"
)

type ListUsersOpts struct { // filters
	ID       string `form:"id"`
	FullName string `form:"full_name"`
	Email    string `form:"email"`
	Phone    string `form:"phone"`
	Role     string `form:"role"`
	Status   string `form:"status"`
	Page     int    `form:"page"`
	Limit    int    `form:"limit"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *Entities.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*Entities.User, error) {
	var user Entities.User
	result := r.db.First(&user, "id = ?", id)
	return &user, result.Error
}



func (r *UserRepository) UpdateUser(user *Entities.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUser(user *Entities.User) error {
	return r.db.Delete(user).Error
}

func (r *UserRepository) ListUsers(opts ListUsersOpts) ([]Entities.User, error) {
	var users []Entities.User
	query := r.db.Model(&Entities.User{}) // Gets users from DB using filters

	if opts.ID != "" {
		query = query.Where("id = ?", opts.ID)
	}

	if opts.FullName != "" {
		name := strings.TrimSpace(opts.FullName)
		query = query.Where("LOWER(full_name) LIKE ?", "%"+strings.ToLower(name)+"%")
	}

	if opts.Email != "" {
		query = query.Where("email = ?", opts.Email)
	}

	if opts.Phone != "" {
		phone := strings.TrimSpace(opts.Phone)
		phone = strings.ReplaceAll(phone, "+", "")
		phone = strings.ReplaceAll(phone, " ", "")
		phone = strings.ReplaceAll(phone, "-", "")
		phone = strings.ReplaceAll(phone, "(", "")
		phone = strings.ReplaceAll(phone, ")", "")
		query = query.Where("REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(phone, '+', ''), ' ', ''), '-', ''), '(', ''), ')', '') = ?", phone)
	}

	if opts.Role != "" {
		query = query.Where("role = ?", opts.Role)
	}

	if opts.Status != "" {
		query = query.Where("status = ?", opts.Status)
	}

	if opts.Limit == 0 {
		opts.Limit = 10
	}
	if opts.Page == 0 {
		opts.Page = 1
	}
	offset := (opts.Page - 1) * opts.Limit

	err := query.
		Limit(opts.Limit).
		Offset(offset).
		Find(&users).Error

	return users, err
}

func (r *UserRepository) GetByEmail(email string) (*Entities.User, error) {
	var user Entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdatePassword(email, hashedPassword string) error {
	user := &Entities.User{}
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}

	user.PasswordHash = hashedPassword
	return r.db.Save(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*Entities.User, error) {
    var user Entities.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}