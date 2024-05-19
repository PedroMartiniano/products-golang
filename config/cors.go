package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsConfig() gin.HandlerFunc {
	configs := cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},
	})

	return configs
}
