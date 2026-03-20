package Entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null;index"`
	TotalAmount float64        `gorm:"not null"`
	Status      string         `gorm:"not null"` 

	OrderItems  []OrderItem    `gorm:"foreignKey:OrderID"`

	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}