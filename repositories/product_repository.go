package repositories

import (
	"go-echo-crud/models"
)

type ProductRepository interface {
	Create(product *models.Product) error
	FindAll() ([]*models.Product, error)
	FindByID(id uint) (*models.Product, error)
	Update(product *models.Product) error
	Delete(product *models.Product) error
}
