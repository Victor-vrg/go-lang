package routes

import (
	"github.com/Victor-vrg/go-lang/controllers"
	"github.com/Victor-vrg/go-lang/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)

	app.Post("/add-task", controllers.AddTask)

	protected := app.Group("/protected")
	protected.Use(middleware.JWTMiddleware)
	protected.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "This is a protected route"})
	})
}
