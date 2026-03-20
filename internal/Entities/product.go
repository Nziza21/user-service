package Entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Name        string         `gorm:"not null"`
	Description string
	Price       float64        `gorm:"not null"`
	Stock       int            `gorm:"not null"`
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}