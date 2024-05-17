package services

import (
	"go-echo-crud/models"
	"go-echo-crud/repositories"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ProductService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	if err := validate.Struct(product); err != nil {
		return err
	}
	return s.repo.Create(product)
}

func (s *ProductService) GetProducts() ([]*models.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) UpdateProduct(id uint, updatedProduct *models.ProductUpdate) (*models.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	updateProductFields(product, updatedProduct)

	if err := validate.Struct(product); err != nil {
		return nil, err
	}
	return product, s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(product)
}

func updateProductFields(product *models.Product, updatedProduct *models.ProductUpdate) {
	v := reflect.ValueOf(updatedProduct).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.IsNil() {
			productField := reflect.ValueOf(product).Elem().FieldByName(t.Field(i).Name)
			if productField.IsValid() {
				productField.Set(reflect.ValueOf(field.Elem().Interface()))
			}
		}
	}
}
