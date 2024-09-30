package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/maksimdudarev/golang-webapi-example/database"
	"github.com/maksimdudarev/golang-webapi-example/models"
)

func GetDiscountList(c *fiber.Ctx) error {
	items := []models.Discount{}

	database.DB.Db.Find(&items)

	return c.Status(fiber.StatusOK).JSON(items)
}

func GetDiscountItem(c *fiber.Ctx) error {
	item := models.Discount{}
	id := c.Params("id")

	database.DB.Db.Find(&item, "id = ?", id)

	if item.ID == uint(uuid.Nil.ID()) {
		return c.Status(fiber.StatusNotFound).JSON(nil)
	}

	return c.Status(fiber.StatusOK).JSON(item)
}

func CreateDiscount(c *fiber.Ctx) error {
	item := new(models.Discount)

	if err := c.BodyParser(item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&item)

	return c.Status(fiber.StatusCreated).JSON(item)
}

func UpdateDiscount(c *fiber.Ctx) error {
	type updateDiscount struct {
		ProductName string
		Description string
		Amount      int
	}

	item := models.Discount{}
	id := c.Params("id")

	database.DB.Db.Find(&item, "id = ?", id)

	if item.ID == uint(uuid.Nil.ID()) {
		return c.Status(fiber.StatusNotFound).JSON(nil)
	}

	var updateItemData updateDiscount
	if err := c.BodyParser(&updateItemData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	item.ProductName = updateItemData.ProductName
	item.Description = updateItemData.Description
	item.Amount = updateItemData.Amount

	err := database.DB.Db.Save(&item).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(id)
}

func DeleteDiscount(c *fiber.Ctx) error {
	item := models.Discount{}
	id := c.Params("id")

	database.DB.Db.Find(&item, "id = ?", id)

	if item.ID == uint(uuid.Nil.ID()) {
		return c.Status(fiber.StatusNotFound).JSON(nil)
	}

	err := database.DB.Db.Delete(&item, "id = ?", id).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(id)
}
