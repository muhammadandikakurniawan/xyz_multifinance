package main

import (
	"github.com/joho/godotenv"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/container"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/delivery"
)

func init() {
	godotenv.Load()
}

func main() {
	dependencyContainer := container.SetupContainer()

	delivery.Run(&dependencyContainer)
}
