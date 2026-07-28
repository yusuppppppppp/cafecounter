package services

import (
	"github.com/yusuppppppppp/cafecounter/apps/api/internal/models"
	"github.com/yusuppppppppp/cafecounter/apps/api/internal/repositories"
)

type productService struct {
	repository repositories.ProductRepository
}

func NewProductService(repository repositories.ProductRepository) ProductService {
	return &productService{
		repository: repository,
	}
}

func (s *productService) Create(product *models.Product) error {
	return s.repository.Create(product)
}

func (s *productService) FindAll() ([]models.Product, error) {
	return s.repository.FindAll()
}

func (s *productService) FindByID(id uint) (*models.Product, error) {
	return s.repository.FindByID(id)
}

func (s *productService) Update(product *models.Product) error {
	return s.repository.Update(product)
}		

func (s *productService) Delete(id uint) error {
	return s.repository.Delete(id)
}
