package repositories

import (
	"github.com/yusuppppppppp/cafecounter/apps/api/internal/models"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category

	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepository) FindByID(id uint) (*models.Category, error) {
	var category models.Category

	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) Update(category *models.Category) error {
	return r.db.Model(category).Updates(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}
