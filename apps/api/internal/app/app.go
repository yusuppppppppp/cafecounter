package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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

	go func() {
		if err := a.server.Listen(a.config.Addr); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-quit

	log.Println("Shutdown signal received...")

	if err := a.server.Shutdown(); err != nil {
		log.Printf("Failed to shutdown server: %v", err)
		return
	}

	log.Println("Server stopped gracefully")
}
