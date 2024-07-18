package services

import (
	"context"

	"github.com/PedroMartiniano/products-golang/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProductService interface {
	CreateProductExecute(context.Context, models.Product) (models.Product, error)
	FindProductByIdExecute(context.Context, primitive.ObjectID) (models.Product, error)
	ListProductsExecute(context.Context) ([]models.Product, error)
	UpdateProductExecute(context.Context, models.Product) (models.Product, error)
	DeleteProductExecute(context.Context, models.Product) (models.Product, error)
}