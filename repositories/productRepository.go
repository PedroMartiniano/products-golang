package repositories

import (
	"context"
	"time"

	"github.com/PedroMartiniano/products-golang/config"
	"github.com/PedroMartiniano/products-golang/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (pr *productRepository) FindById(id primitive.ObjectID) (models.Product, error) {
	collection := config.DB.Collection("products")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var product models.Product
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (pr *productRepository) List() ([]models.Product, error) {
	collection := config.DB.Collection("products")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []models.Product{}, err
	}

	var products []models.Product
	err = cur.All(ctx, &products)
	if err != nil {
		return []models.Product{}, err
	}

	return products, nil
}

func (pr *productRepository) Update(product models.Product) (models.Product, error) {
	collection := config.DB.Collection("products")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	updatedProduct := bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"category":    product.Category,
			"updated_at":  time.Now(),
		},
	}

	_, err := collection.UpdateByID(ctx, product.ID, updatedProduct)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (pr *productRepository) Delete(product models.Product) (models.Product, error) {
	collection := config.DB.Collection("products")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": product.ID})
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}
