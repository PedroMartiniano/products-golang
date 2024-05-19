package services

import (
	"github.com/PedroMartiniano/products-golang/models"
	ports "github.com/PedroMartiniano/products-golang/ports/repositories"
)

type productService struct {
	productRepository ports.IProductRepository
}

func NewProductService(productRepository ports.IProductRepository) *productService {
	return &productService{
		productRepository: productRepository,
	}
}

func (ps *productService) CreateProductExecute(product models.Product) (models.Product, error) {
	newProduct, err := ps.productRepository.Create(product)

	return newProduct, err
}
