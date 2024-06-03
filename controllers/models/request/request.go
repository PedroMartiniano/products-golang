package request

type createProductRequest struct {
	Name        string  `json:"name" bson:"name" binding:"required"`
	Description string  `json:"description" bson:"description" binding:"required"`
	Price       float64 `json:"price" bson:"price" binding:"required"`
	Category    string  `json:"category" bson:"category" binding:"required"`
}

type updateProductRequest struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
	Category    string  `json:"category" bson:"category"`
}
