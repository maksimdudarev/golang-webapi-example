package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maksimdudarev/golang-webapi-example/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})

	discount := app.Group("/discount")
	discount.Get("/", handlers.GetDiscountList)
	discount.Get("/:id", handlers.GetDiscountItem)
	discount.Post("/", handlers.CreateDiscount)
	discount.Put("/:id", handlers.UpdateDiscount)
	discount.Delete("/:id", handlers.DeleteDiscount)
}
