package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maksimdudarev/golang-webapi-example/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})

	fact := app.Group("/fact")
	fact.Get("/", handlers.GetFactList)
	fact.Get("/:id", handlers.GetFactItem)
	fact.Post("/", handlers.CreateFact)
	fact.Put("/:id", handlers.UpdateFact)
	fact.Delete("/:id", handlers.DeleteFact)
}
