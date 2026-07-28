package services

import (
	"github.com/yusuppppppppp/cafecounter/apps/api/internal/models"
	"github.com/yusuppppppppp/cafecounter/apps/api/internal/repositories"
)

type categoryService struct {
	repository repositories.CategoryRepository
}

func NewCategoryService(repository repositories.CategoryRepository) CategoryService {
	return &categoryService{
		repository: repository,
	}
}

func (s *categoryService) Create(category *models.Category) error {
	return s.repository.Create(category)
}

func (s *categoryService) FindAll() ([]models.Category, error) {
	return s.repository.FindAll()
}

func (s *categoryService) FindByID(id uint) (*models.Category, error) {
	return s.repository.FindByID(id)
}

func (s *categoryService) Update(category *models.Category) error {
	return s.repository.Update(category)
}

func (s *categoryService) Delete(id uint) error {
	return s.repository.Delete(id)
}
