package controllers

import (
	"mithril/src/database"
	"mithril/src/models"

	"github.com/gofiber/fiber/v2"
)

func Mithril(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Where("is_mithril = true").Find(&users)

	return c.JSON(users)
}
