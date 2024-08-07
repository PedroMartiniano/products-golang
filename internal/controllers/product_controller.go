package controllers

import (
	"fmt"

	"github.com/PedroMartiniano/products-golang/internal/models"
	ps "github.com/PedroMartiniano/products-golang/internal/ports/services"
	"github.com/PedroMartiniano/products-golang/internal/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	productService ps.IProductService
}

func NewProductController(productService ps.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
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
func (pc *ProductController) CreateProductHandler(c *gin.Context) {
	var request createProductRequest

	err := c.BindJSON(&request)
	if err != nil {
		sendError(c, 400, err.Error())
		return
	}

	product := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Weight:      request.Weight,
		Color:       request.Color,
		Type:        request.Type,
	}

	newProduct, err := pc.productService.CreateProductExecute(c.Request.Context(), product)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
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
func (pc *ProductController) FindProductByIdHandler(c *gin.Context) {
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

	product, err := pc.productService.FindProductByIdExecute(c.Request.Context(), objectId)

	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
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
func (pc *ProductController) ListProductsHandler(c *gin.Context) {
	products, err := pc.productService.ListProductsExecute(c.Request.Context())
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, 200, products)
}

// @BasePath /product
// @Summary      Update a product
// @Description  Should update an existing product in the database successfully
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Id Product"
// @Param        request   body      updateProductRequest  true  "Body Json"
// @Success      200  {object}  productSuccessResponse1
// @Failure      400  {object}  errorResponse
// @Router       /product/{id} [put]
func (pc *ProductController) UpdateProductHandler(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		sendError(c, 400, "missing id param")
		return
	}

	var request updateProductRequest

	err := c.BindJSON(&request)
	if err != nil {
		sendError(c, 400, err.Error())
		return
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		sendError(c, 400, err.Error())
		return
	}

	product, err := pc.productService.FindProductByIdExecute(c.Request.Context(), objectId)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	utils.ReplaceStrIfNotEmpty(&product.Name, request.Name)
	utils.ReplaceStrIfNotEmpty(&product.Description, request.Description)
	utils.ReplaceNumIfNotEmpty(&product.Price, request.Price)
	utils.ReplaceNumIfNotEmpty(&product.Weight, request.Weight)
	utils.ReplaceStrIfNotEmpty(&product.Color, request.Color)
	utils.ReplaceStrIfNotEmpty(&product.Type, request.Type)

	fmt.Println(product)
	updatedProduct, err := pc.productService.UpdateProductExecute(c.Request.Context(), product)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, 200, updatedProduct)
}

// @BasePath /product
// @Summary      Delete a product
// @Description  Should delete a product by his id successfully
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Id Product"
// @Success      200  {object}  productSuccessResponse1
// @Failure      400  {object}  errorResponse
// @Router       /product/{id} [delete]
func (pc *ProductController) DeleteProductHandler(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		sendError(c, 400, "missing id param")
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		sendError(c, 400, err.Error())
		return
	}

	product, err := pc.productService.FindProductByIdExecute(c.Request.Context(), objectId)

	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	deletedProduct, err := pc.productService.DeleteProductExecute(c.Request.Context(), product)
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, 200, deletedProduct)
}
