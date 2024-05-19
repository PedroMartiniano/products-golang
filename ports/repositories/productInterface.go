package ports

import "github.com/PedroMartiniano/products-golang/models"

type IProductRepository interface {
	Create(models.Product) (models.Product, error)
	// FindById()
	// List()
	// Update()
	// Delete()
}
