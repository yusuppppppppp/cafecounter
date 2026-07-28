package repositories

import (
	"github.com/yusuppppppppp/cafecounter/apps/api/internal/models"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{	
		db: db,
	}
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product

	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product

	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Model(product).Updates(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
