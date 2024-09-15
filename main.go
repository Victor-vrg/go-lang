package main

import (
	"log"

	"github.com/Victor-vrg/go-lang/config"
	"github.com/Victor-vrg/go-lang/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDB()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", config.DB)
		return c.Next()
	})

	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
