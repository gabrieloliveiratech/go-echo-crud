package repositories

import (
	"errors"
	"go-echo-crud/models"
)

type mockProductRepository struct {
	products map[uint]*models.Product
	nextID   uint
}

func NewMockProductRepository() ProductRepository {
	return &mockProductRepository{
		products: make(map[uint]*models.Product),
		nextID:   1,
	}
}

func (r *mockProductRepository) Create(product *models.Product) error {
	product.ID = r.nextID
	r.products[r.nextID] = product
	r.nextID++
	return nil
}

func (r *mockProductRepository) FindAll() ([]*models.Product, error) {
	var products []*models.Product
	for _, product := range r.products {
		products = append(products, product)
	}
	return products, nil
}

func (r *mockProductRepository) FindByID(id uint) (*models.Product, error) {
	product, exists := r.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (r *mockProductRepository) Update(product *models.Product) error {
	_, exists := r.products[product.ID]
	if !exists {
		return errors.New("product not found")
	}
	r.products[product.ID] = product
	return nil
}

func (r *mockProductRepository) Delete(product *models.Product) error {
	_, exists := r.products[product.ID]
	if !exists {
		return errors.New("product not found")
	}
	delete(r.products, product.ID)
	return nil
}
