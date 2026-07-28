package repositories

import (
	"github.com/yusuppppppppp/cafecounter/apps/api/internal/models"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	Update(category *models.Category) error
	Delete(id uint) error
	FindAll() ([]models.Category, error)
	FindByID(id uint) (*models.Category, error)
}

type ProductRepository interface {
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id uint) error
	FindAll() ([]models.Product, error)
	FindByID(id uint) (*models.Product, error)
}
