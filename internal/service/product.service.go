package service

import (
    "github.com/Nziza21/user-service/internal/Entities"
    "github.com/Nziza21/user-service/internal/store/repository"
)

type ProductService struct {
    repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(p *Entities.Product) error {
    return s.repo.CreateProduct(p)
}

func (s *ProductService) ListProducts() ([]Entities.Product, error) {
    return s.repo.ListProducts()
}

// You can add UpdateProduct and DeleteProduct later similarly