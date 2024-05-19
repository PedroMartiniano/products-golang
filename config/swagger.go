package config

import "github.com/swaggo/swag"

func SwaggerConfig(config *swag.Spec) {
	config.Title = "API Products"
	config.Description = "Documentação pública de uma api de product feita em golang"
	config.Version = "1.0.0"
	config.Host = "localhost:8080"
	config.BasePath = "/"
	config.Schemes = []string{"http", "https"}
}
