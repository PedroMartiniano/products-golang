package routes

import (
	"github.com/PedroMartiniano/products-golang/internal/adapters"
	"github.com/PedroMartiniano/products-golang/internal/controllers"
	"github.com/gin-gonic/gin"
)

func productRoutes(router *gin.RouterGroup) {
	productService := adapters.NewProductServiceAdapter()
	productController := controllers.NewProductController(productService)

	router.POST("/", productController.CreateProductHandler)
	router.GET("/", productController.ListProductsHandler)
	router.GET("/:id", productController.FindProductByIdHandler)
	router.PUT("/:id", productController.UpdateProductHandler)
	router.DELETE("/:id", productController.DeleteProductHandler)
}
