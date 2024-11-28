package main

import (
	"be-hire-revamp/src/config"
	"be-hire-revamp/src/helper"
	"be-hire-revamp/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helper.Migrate()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! ini yohanes")
	})

	routes.Router(app)
	app.Listen(":8080")
}
