package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maksimdudarev/golang-webapi-example/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Post("/fact", handlers.CreateFact)
}
