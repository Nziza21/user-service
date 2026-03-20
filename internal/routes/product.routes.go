package routes

import (
	"github.com/Nziza21/user-service/handler"
	"github.com/gin-gonic/gin"
)

func setupProductRoutes(api *gin.RouterGroup, h *handler.ProductHandler) {
    products := api.Group("/products")  
    products.POST("", h.CreateProduct)
    products.GET("", h.ListProducts)
}