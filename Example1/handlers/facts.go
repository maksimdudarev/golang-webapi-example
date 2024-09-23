package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/maksimdudarev/golang-webapi-example/database"
	"github.com/maksimdudarev/golang-webapi-example/models"
)

func GetFactList(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Status(fiber.StatusOK).JSON(facts)
}

func GetFactItem(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	database.DB.Db.Find(&fact, "id = ?", id)

	if fact.ID == uint(uuid.Nil.ID()) {
		return c.Status(fiber.StatusNotFound).JSON(nil)
	}

	return c.Status(fiber.StatusOK).JSON(fact)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(fiber.StatusCreated).JSON(fact)
}

func UpdateFact(c *fiber.Ctx) error {
	type updateFact struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}

	fact := models.Fact{}
	id := c.Params("id")

	database.DB.Db.Find(&fact, "id = ?", id)

	if fact.ID == uint(uuid.Nil.ID()) {
		return c.Status(fiber.StatusNotFound).JSON(nil)
	}

	var updateFactData updateFact
	if err := c.BodyParser(&updateFactData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fact.Question = updateFactData.Question
	fact.Answer = updateFactData.Answer

	err := database.DB.Db.Save(&fact).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(id)
}

func DeleteFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	database.DB.Db.Find(&fact, "id = ?", id)

	if fact.ID == uint(uuid.Nil.ID()) {
		return c.Status(fiber.StatusNotFound).JSON(nil)
	}

	err := database.DB.Db.Delete(&fact, "id = ?", id).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(id)
}
