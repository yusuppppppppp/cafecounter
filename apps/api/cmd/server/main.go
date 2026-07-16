package main

import (
	"github.com/yusuppppppppp/cafecounter/apps/api/configs"
	"github.com/yusuppppppppp/cafecounter/apps/api/internal/app"
)

func main() {
	cfg := configs.Load()

	application := app.New(cfg)

	application.Start()
}
