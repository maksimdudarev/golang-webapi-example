package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maksimdudarev/golang-webapi-example/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
