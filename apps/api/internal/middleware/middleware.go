package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Register(app *fiber.App) {
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "15:04:05",
		TimeZone:   "Asia/Jakarta",
		Format:     "[${time}] ${status} | ${latency} | ${method} ${path}\n",
	}))
}
