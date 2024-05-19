package routes

import (
	"github.com/PedroMartiniano/products-golang/controllers"
	"github.com/gin-gonic/gin"
)

func productRoutes(router *gin.RouterGroup) {
	productController := controllers.NewProductController()

	router.POST("/", productController.CreateProductHandler)
}
