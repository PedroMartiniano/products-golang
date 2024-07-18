package ports

import (
	"context"

	"github.com/PedroMartiniano/products-golang/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProductRepository interface {
	Create(context.Context, models.Product) (models.Product, error)
	FindById(context.Context, primitive.ObjectID) (models.Product, error)
	List(context.Context) ([]models.Product, error)
	Update(context.Context, models.Product) (models.Product, error)
	Delete(context.Context, models.Product) (models.Product, error)
}
