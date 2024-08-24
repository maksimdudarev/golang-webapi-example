package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maksimdudarev/golang-webapi-example/database"
	"github.com/maksimdudarev/golang-webapi-example/models"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Status(fiber.StatusOK).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(fiber.StatusOK).JSON(fact)
}
