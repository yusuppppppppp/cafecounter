package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yusuppppppppp/cafecounter/apps/api/configs"
	"github.com/yusuppppppppp/cafecounter/apps/api/internal/middleware"
	"github.com/yusuppppppppp/cafecounter/apps/api/internal/routes"
)

type App struct {
	server *fiber.App
	config *configs.Config
}

func New(cfg *configs.Config) *App {
	server := fiber.New()

	middleware.Register(server)
	routes.Register(server)

	return &App{
		server: server,
		config: cfg,
	}
}

func (a *App) Start() {
	log.Printf("%s running on http://%s", a.config.AppName, a.config.Addr)

	log.Fatal(a.server.Listen(a.config.Addr))
}
