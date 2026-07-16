package app

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/yusuppppppppp/cafecounter/apps/api/configs"
)

type App struct {
	server *fiber.App
	config *configs.Config
}

func New(cfg *configs.Config) *App {
	server := fiber.New()

	return &App{
		server: server,
		config: cfg,
	}
}

func (a *App) Start() {
	log.Printf("%s running on http://%s", a.config.AppName, a.config.Addr)

	log.Fatal(a.server.Listen(a.config.Addr))
}
