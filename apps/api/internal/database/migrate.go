package database

import (
	"github.com/yusuppppppppp/cafecounter/apps/api/internal/models"
)

func (d *Database) Migrate() error {
	return d.DB.AutoMigrate(
		&models.Category{},
		&models.Product{},
	)
}
