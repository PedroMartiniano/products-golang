package main

import (
	"fmt"

	"github.com/PedroMartiniano/products-golang/internal/config"
	"github.com/PedroMartiniano/products-golang/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	gin.SetMode(gin.DebugMode)

	server := gin.Default()

	server.Use(config.CorsConfig())

	routes.RegisterRoutes(server)

	port := config.GetEnv("PORT")
	server.Run(fmt.Sprintf(":%s", port))
}
