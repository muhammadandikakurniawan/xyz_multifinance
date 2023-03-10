package main

import (
	"github.com/joho/godotenv"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/container"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/delivery"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	dependencyContainer := container.SetupContainer()

	delivery.Run(&dependencyContainer)
}
