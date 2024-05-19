package ports

import (
	"github.com/PedroMartiniano/products-golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProductRepository interface {
	Create(models.Product) (models.Product, error)
	FindById(id primitive.ObjectID) (models.Product, error)
	List() ([]models.Product, error)
	// Update()
	// Delete()
}
