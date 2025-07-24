package main

import "github.com/Amir-Sadati/crypto-microservice/crypto-order-service/internal/app"

func main() {
	app := app.New()
	app.Run()
}
