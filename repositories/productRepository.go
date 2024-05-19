package repositories

import (
	"context"
	"time"

	"github.com/PedroMartiniano/products-golang/config"
	"github.com/PedroMartiniano/products-golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productRepository struct{}

func NewProductRepository() *productRepository {
	return &productRepository{}
}

func (pr *productRepository) Create(product models.Product) (models.Product, error) {
	collection := config.DB.Collection("products")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}
