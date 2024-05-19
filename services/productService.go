package services

import (
	"github.com/PedroMartiniano/products-golang/models"
	ports "github.com/PedroMartiniano/products-golang/ports/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (ps *productService) FindProductByIdExecute(id primitive.ObjectID) (models.Product, error) {
	product, err := ps.productRepository.FindById(id)

	return product, err
}

func (ps *productService) ListProductsExecute() ([]models.Product, error) {
	products, err := ps.productRepository.List()

	return products, err
}
