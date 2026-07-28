package services

import "github.com/yusuppppppppp/cafecounter/apps/api/internal/models"

type ProductService interface {
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id uint) error
	FindAll() ([]models.Product, error)
	FindByID(id uint) (*models.Product, error)
}