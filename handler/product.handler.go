package handler

import (
    "net/http"
    "github.com/Nziza21/user-service/internal/Entities"
    "github.com/Nziza21/user-service/internal/service"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

type ProductHandler struct {
    productService *service.ProductService
}

type ErrorResponse struct {
    Error string `json:"error"`
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
    return &ProductHandler{productService: s}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a product with name, description, price, and stock
// @Tags products
// @Accept json
// @Produce json
// @Param product body CreateProductRequest true "Product Info"
// @Success 201 {object} ProductResponse
// @Failure 400 {object} map[string]string
// @Router /api/v1/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
    // 1. Bind input
    var req CreateProductRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 2. Create Product object for DB
    p := Entities.Product{
        ID:          uuid.New(),
        Name:        req.Name,
        Description: req.Description,
        Price:       req.Price,
        Stock:       req.Stock,
    }

    // 3. Save to DB
    if err := h.productService.CreateProduct(&p); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 4. Return response
    res := ProductResponse{
        ID:          p.ID.String(),
        Name:        p.Name,
        Description: p.Description,
        Price:       p.Price,
        Stock:       p.Stock,
    }

    c.JSON(http.StatusCreated, res)
}

// ListProducts godoc
// @Summary List all products
// @Description Get a list of all products
// @Tags products
// @Produce json
// @Success 200 {array} handler.ProductResponse
// @Router /api/v1/products [get]
func (h *ProductHandler) ListProducts(c *gin.Context) {
    products, err := h.productService.ListProducts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var res []ProductResponse
    for _, p := range products {
        res = append(res, ProductResponse{
            ID:          p.ID.String(),
            Name:        p.Name,
            Description: p.Description,
            Price:       p.Price,
            Stock:       p.Stock,
        })
    }

    c.JSON(http.StatusOK, res)
}