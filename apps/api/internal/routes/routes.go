package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("CafeCounter API is running")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"message": "CafeCounter API is healthy",
		})
	})
}
