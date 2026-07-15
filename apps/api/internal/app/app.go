package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

const Port = ":8080"

type App struct {
	server *fiber.App
}

func New() *App {
	server := fiber.New()

	return &App{
		server: server,
	}
}

func (a *App) Start() {
	log.Printf("CafeCounter Api running on http://localhost%s", Port)

	log.Fatal(a.server.Listen(Port))
}
