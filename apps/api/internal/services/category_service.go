package services

import "github.com/yusuppppppppp/cafecounter/apps/api/internal/models"

type CategoryService interface {
	Create(category *models.Category) error
	Update(category *models.Category) error
	Delete(id uint) error
	FindAll() ([]models.Category, error)
	FindByID(id uint) (*models.Category, error)
}