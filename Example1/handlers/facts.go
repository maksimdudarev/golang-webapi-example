package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maksimdudarev/golang-webapi-example/database"
	"github.com/maksimdudarev/golang-webapi-example/models"
)

func GetFactList(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Status(fiber.StatusOK).JSON(facts)
}

func GetFactItem(c *fiber.Ctx) error {
	return c.SendString("I'm a GET item request!")
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

func UpdateFact(c *fiber.Ctx) error {
	return c.SendString("I'm a PUT request!")
}

func DeleteFact(c *fiber.Ctx) error {
	return c.SendString("I'm a DELETE request!")
}
