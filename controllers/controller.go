package controllers

import (
	"apiant/models"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var ant []models.Antri
	models.DB.Db.Find(&ant)
	return c.Status(fiber.StatusOK).JSON(ant)
}

func Show(c *fiber.Ctx) error {
	ant := &models.Antri{}
	id := c.Params("id")
	if err := models.DB.Db.First(ant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(ant)
}

func Update(c *fiber.Ctx) error {
	ant := &models.Antri{}
	id := c.Params("id")
	if err := models.DB.Db.First(ant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	if err := c.BodyParser(ant); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	ant.Seq = ant.Seq + 1
	models.DB.Db.Save(ant)
	return c.Status(fiber.StatusOK).JSON(ant)
}

func Reset(c *fiber.Ctx) error {
	ant := &models.Antri{}
	id := c.Params("id")
	if err := models.DB.Db.First(ant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}
	if err := c.BodyParser(ant); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	ant.Seq = 0
	models.DB.Db.Save(ant)
	return c.Status(fiber.StatusOK).JSON(ant)
}
