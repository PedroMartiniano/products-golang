package controllers

import (
	"github.com/PedroMartiniano/products-golang/models"
	"github.com/PedroMartiniano/products-golang/repositories"
	"github.com/PedroMartiniano/products-golang/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productController struct{}

func NewProductController() *productController {
	return &productController{}
}

// @BasePath /product
// @Summary      Create a new product
// @Description  Should create a new product in the database successfully
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        request   body      createProductRequest  true  "Body Json"
// @Success      201  {object}  productSuccessResponse1
// @Failure      400  {object}  errorResponse
// @Router       /product [post]
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

// @BasePath /product
// @Summary      Get a product by his id
// @Description  Should get a product by his id successfully
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Id Product"
// @Success      200  {object}  productSuccessResponse1
// @Failure      400  {object}  errorResponse
// @Router       /product/{id} [get]
func (pc *productController) FindProductByIdHandler(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		sendError(c, 400, "missing id param")
		return
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		sendError(c, 400, err.Error())
		return
	}

	productRepository := repositories.NewProductRepository()
	productService := services.NewProductService(productRepository)

	product, err := productService.FindProductByIdExecute(objectId)

	if err != nil {
		sendError(c, 400, err.Error())
		return
	}

	sendSuccess(c, 200, product)
}

// @BasePath /product
// @Summary      Get all products
// @Description  Should get all products successfully
// @Tags         Product
// @Accept       json
// @Produce      json
// @Success      200  {object}  productSuccessResponse2
// @Failure      400  {object}  errorResponse
// @Router       /product [get]
func (pc *productController) ListProductsHandler(c *gin.Context) {
	productRepository := repositories.NewProductRepository()
	productService := services.NewProductService(productRepository)

	products, err := productService.ListProductsExecute()

	if err != nil {
		sendError(c, 400, err.Error())
		return
	}

	sendSuccess(c, 200, products)
}
