package main

import (
	"mithril/src/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectWithAutoMigrate()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World !")
	})

	app.Listen(":8000")
}
