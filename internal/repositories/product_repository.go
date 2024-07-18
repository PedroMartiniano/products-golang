package repositories

import (
	"context"
	"time"

	"github.com/PedroMartiniano/products-golang/internal/config"
	"github.com/PedroMartiniano/products-golang/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (pr *ProductRepository) Create(c context.Context, product models.Product) (models.Product, error) {
	collection := config.DB.Collection("products")

	ctx, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		return models.Product{}, config.NewError(config.ErrInternalServer, err)
	}

	return product, nil
}

func (pr *ProductRepository) FindById(c context.Context, id primitive.ObjectID) (models.Product, error) {
	collection := config.DB.Collection("products")

	ctx, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	var product models.Product
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return models.Product{}, config.NewError(config.ErrNotFound, err)
	}

	return product, nil
}

func (pr *ProductRepository) List(c context.Context) ([]models.Product, error) {
	collection := config.DB.Collection("products")

	ctx, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return []models.Product{}, config.NewError(config.ErrInternalServer, err)
	}

	var products []models.Product
	err = cur.All(ctx, &products)
	if err != nil {
		return []models.Product{}, config.NewError(config.ErrInternalServer, err)
	}

	return products, nil
}

func (pr *ProductRepository) Update(c context.Context, product models.Product) (models.Product, error) {
	collection := config.DB.Collection("products")

	ctx, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	updatedProduct := bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"weight":      product.Weight,
			"color":       product.Color,
			"type":        product.Type,
			"updated_at":  time.Now(),
		},
	}

	_, err := collection.UpdateByID(ctx, product.ID, updatedProduct)
	if err != nil {
		return models.Product{}, config.NewError(config.ErrInternalServer, err)
	}

	return product, nil
}

func (pr *ProductRepository) Delete(c context.Context, product models.Product) (models.Product, error) {
	collection := config.DB.Collection("products")

	ctx, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": product.ID})
	if err != nil {
		return models.Product{}, config.NewError(config.ErrInternalServer, err)
	}

	return product, nil
}
