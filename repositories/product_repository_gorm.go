package repositories

import (
	"go-echo-crud/config"
	"go-echo-crud/models"

	"gorm.io/gorm"
)

type gormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository() ProductRepository {
	return &gormProductRepository{
		db: config.DB, // Assumindo que `config.DB` é a conexão com o banco de dados
	}
}

func (r *gormProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *gormProductRepository) FindAll() ([]*models.Product, error) {
	var products []*models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *gormProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *gormProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *gormProductRepository) Delete(product *models.Product) error {
	return r.db.Delete(product).Error
}
