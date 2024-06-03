package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Color       string             `json:"color" bson:"color"`
	Weight      float64            `json:"weight" bson:"weight"`
	Type        string             `json:"type" bson:"type"`
	Price       float64            `json:"price" bson:"price"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
