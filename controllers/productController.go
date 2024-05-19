package controllers

import (
	"github.com/PedroMartiniano/products-golang/models"
	"github.com/PedroMartiniano/products-golang/repositories"
	"github.com/PedroMartiniano/products-golang/services"
	"github.com/gin-gonic/gin"
)

type productController struct{}

func NewProductController() *productController {
	return &productController{}
}

func (pc *productController) CreateProductHandler(c *gin.Context) {
	var request createProductRequest

	err := c.BindJSON(&request)
	if err != nil {
		sendError(c, 400, err.Error())
		return
	}

	product := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Category:    request.Category,
		Price:       request.Price,
	}

	productRepository := repositories.NewProductRepository()
	productService := services.NewProductService(productRepository)

	newProduct, err := productService.CreateProductExecute(product)
	if err != nil {
		sendError(c, 400, err.Error())
		return
	}

	sendSuccess(c, 201, newProduct)
}
