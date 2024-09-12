package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maksimdudarev/golang-webapi-example/database"
	"github.com/maksimdudarev/golang-webapi-example/routes"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
