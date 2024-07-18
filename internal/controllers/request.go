package controllers

type createProductRequest struct {
	Name        string  `json:"name" bson:"name" binding:"required"`
	Description string  `json:"description" bson:"description" binding:"required"`
	Price       float64 `json:"price" bson:"price" binding:"required"`
	Color       string  `json:"color" bson:"color" binding:"required"`
	Weight      float64 `json:"weight" bson:"weight" binding:"required"`
	Type        string  `json:"type" bson:"type" binding:"required"`
}

type updateProductRequest struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
	Color       string  `json:"color" bson:"color"`
	Weight      float64 `json:"Weight" bson:"Weight"`
	Type        string  `json:"type" bson:"type"`
}
