package config

import "github.com/swaggo/swag"

func SwaggerConfig(config *swag.Spec) {
	deploy_url := GetEnv("DEPLOY_URL")

	config.Title = "API Products"
	config.Description = "Documentação pública da api de products feita em Golang por Pedro Paulino Martiniano"
	config.Version = "1.0.0"
	config.Host = deploy_url
	config.BasePath = "/"
	config.Schemes = []string{"https", "http"}
}
