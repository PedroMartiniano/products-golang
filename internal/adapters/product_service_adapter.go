package adapters

import (
	"github.com/PedroMartiniano/products-golang/internal/repositories"
	"github.com/PedroMartiniano/products-golang/internal/services"
	ps "github.com/PedroMartiniano/products-golang/internal/ports/services"
)

func NewProductServiceAdapter() ps.IProductService {
	productRepository := repositories.NewProductRepository()
	productService := services.NewProductService(productRepository)

	return productService
}
