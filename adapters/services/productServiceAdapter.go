package adapters

import (
	"github.com/PedroMartiniano/products-golang/repositories"
	"github.com/PedroMartiniano/products-golang/services"
)

func NewProductServiceAdapter() *services.ProductService {
	productRepository := repositories.NewProductRepository()
	productService := services.NewProductService(productRepository)

	return productService
}
