package Entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	FullName     string         `gorm:"size:120;not null" json:"fullName" binding:"required,min=3"`
	Email        string         `gorm:"uniqueIndex;not null" json:"email" binding:"required,email"`
	Phone        string         `json:"phone,omitempty" binding:"omitempty,regexp=^\\+?[0-9]{10,15}$"`
	PasswordHash string         `gorm:"type:text" json:"-"`
	Role         string         `gorm:"size:32;default:'user'" json:"role" binding:"omitempty,oneof=user admin"`
	Status       string         `gorm:"size:16;default:'active'" json:"status" binding:"omitempty,oneof=active inactive"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}