package services

import (
	"github.com/PedroMartiniano/products-golang/models"
	ports "github.com/PedroMartiniano/products-golang/ports/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct {
	productRepository ports.IProductRepository
}

func NewProductService(productRepository ports.IProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (ps *ProductService) CreateProductExecute(product models.Product) (models.Product, error) {
	newProduct, err := ps.productRepository.Create(product)

	return newProduct, err
}

func (ps *ProductService) FindProductByIdExecute(id primitive.ObjectID) (models.Product, error) {
	product, err := ps.productRepository.FindById(id)

	return product, err
}

func (ps *ProductService) ListProductsExecute() ([]models.Product, error) {
	products, err := ps.productRepository.List()

	return products, err
}

func (ps *ProductService) UpdateProductExecute(product models.Product) (models.Product, error) {
	updatedProduct, err := ps.productRepository.Update(product)

	return updatedProduct, err
}

func (ps *ProductService) DeleteProductExecute(product models.Product) (models.Product, error) {
	productDeleted, err := ps.productRepository.Delete(product)

	return productDeleted, err
}
