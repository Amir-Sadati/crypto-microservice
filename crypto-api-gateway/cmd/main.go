package main

import (
	"github.com/Amir-Sadati/crypto-microservice/crypto-api-gateway/internal/app"
)

func main() {
	app := app.New()
	app.Run()
}
