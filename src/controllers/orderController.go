package controllers

import (
	"mithril/src/database"
	"mithril/src/models"

	"github.com/gofiber/fiber/v2"
)

func Orders(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Preload("OrderItems").Find(&orders)

	for i, order := range orders {
		orders[i].Name = order.FullName()
		orders[i].Total = order.TotalValue()
	}

	return c.JSON(orders)
}
