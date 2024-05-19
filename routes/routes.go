package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	productRouter := server.Group("/product")
	productRoutes(productRouter)
}
