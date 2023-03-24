package main

import (
	"mithril/src/database"
	"mithril/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectWithAutoMigrate()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
