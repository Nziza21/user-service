package repository

import (
    "github.com/Nziza21/user-service/internal/Entities"
    "gorm.io/gorm"
)

type ProductRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
    return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(p *Entities.Product) error {
    return r.db.Create(p).Error
}

func (r *ProductRepository) ListProducts() ([]Entities.Product, error) {
    var products []Entities.Product
    err := r.db.Find(&products).Error
    return products, err
}