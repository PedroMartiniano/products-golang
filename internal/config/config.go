package config

import (
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Database

func Init() {
	godotenv.Load()

	var err error

	DB, err = InitMongoDB()
	if err != nil {
		panic(err)
	}
}
