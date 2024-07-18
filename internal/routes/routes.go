package routes

import (
	"github.com/PedroMartiniano/products-golang/internal/config"
	"github.com/PedroMartiniano/products-golang/internal/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(server *gin.Engine) {
	config.SwaggerConfig(docs.SwaggerInfo)

	productRouter := server.Group("/product")
	productRoutes(productRouter)

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
