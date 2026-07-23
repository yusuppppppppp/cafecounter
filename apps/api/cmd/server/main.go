package main

import (
	"log"

	"github.com/yusuppppppppp/cafecounter/apps/api/configs"
	"github.com/yusuppppppppp/cafecounter/apps/api/internal/app"
	"github.com/yusuppppppppp/cafecounter/apps/api/internal/database"
)

func main() {
	cfg := configs.Load()

	db := database.New()

	if err := db.Connect(cfg.DataBaseURL); err != nil {
		log.Fatal(err)
	}

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	application := app.New(cfg, db)

	application.Start()
}
