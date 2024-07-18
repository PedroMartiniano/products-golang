package services

import (
	"context"

	"github.com/PedroMartiniano/products-golang/internal/models"
	pr "github.com/PedroMartiniano/products-golang/internal/ports/repositories"
	ps "github.com/PedroMartiniano/products-golang/internal/ports/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct {
	productRepository pr.IProductRepository
}

func NewProductService(productRepository pr.IProductRepository) ps.IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (ps *ProductService) CreateProductExecute(c context.Context, product models.Product) (models.Product, error) {
	newProduct, err := ps.productRepository.Create(c, product)

	return newProduct, err
}

func (ps *ProductService) FindProductByIdExecute(c context.Context, id primitive.ObjectID) (models.Product, error) {
	product, err := ps.productRepository.FindById(c, id)

	return product, err
}

func (ps *ProductService) ListProductsExecute(c context.Context) ([]models.Product, error) {
	products, err := ps.productRepository.List(c)

	return products, err
}

func (ps *ProductService) UpdateProductExecute(c context.Context, product models.Product) (models.Product, error) {
	updatedProduct, err := ps.productRepository.Update(c, product)

	return updatedProduct, err
}

func (ps *ProductService) DeleteProductExecute(c context.Context, product models.Product) (models.Product, error) {
	productDeleted, err := ps.productRepository.Delete(c, product)

	return productDeleted, err
}
