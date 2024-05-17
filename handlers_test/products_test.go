package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"go-echo-crud/handlers"
	"go-echo-crud/models"
	"go-echo-crud/repositories"
	"go-echo-crud/services"
)

func TestCreateProduct(t *testing.T) {
	mockRepo := repositories.NewMockProductRepository()
	mockService := services.NewProductService(mockRepo)
	handler := handlers.NewProductHandler(mockService)

	product := &models.Product{Name: "Test Product", Price: 100, Stock: 10}
	productJSON, _ := json.Marshal(product)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/products", bytes.NewReader(productJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.CreateProduct(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var newProduct models.Product
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &newProduct)) {
			assert.Equal(t, product.Name, newProduct.Name)
			assert.Equal(t, product.Price, newProduct.Price)
			assert.Equal(t, product.Stock, newProduct.Stock)
		}
	}
}

func TestGetProducts(t *testing.T) {
	mockRepo := repositories.NewMockProductRepository()
	mockService := services.NewProductService(mockRepo)
	handler := handlers.NewProductHandler(mockService)

	product := &models.Product{Name: "Test Product", Price: 100, Stock: 10}
	mockRepo.Create(product)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.GetProducts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var products []models.Product
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &products)) {
			assert.Len(t, products, 1)
			assert.Equal(t, product.Name, products[0].Name)
			assert.Equal(t, product.Price, products[0].Price)
			assert.Equal(t, product.Stock, products[0].Stock)
		}
	}
}

func TestGetProduct(t *testing.T) {
	mockRepo := repositories.NewMockProductRepository()
	mockService := services.NewProductService(mockRepo)
	handler := handlers.NewProductHandler(mockService)

	product := &models.Product{Name: "Test Product", Price: 100, Stock: 10}
	mockRepo.Create(product)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/products/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(product.ID)))

	if assert.NoError(t, handler.GetProduct(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var returnedProduct models.Product
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &returnedProduct)) {
			assert.Equal(t, product.Name, returnedProduct.Name)
			assert.Equal(t, product.Price, returnedProduct.Price)
			assert.Equal(t, product.Stock, returnedProduct.Stock)
		}
	}
}

func TestUpdateProduct(t *testing.T) {
	mockRepo := repositories.NewMockProductRepository()
	mockService := services.NewProductService(mockRepo)
	handler := handlers.NewProductHandler(mockService)

	product := &models.Product{Name: "Test Product", Price: 100, Stock: 10}
	mockRepo.Create(product)
	updatedProduct := &models.Product{Name: "Updated Product", Price: 200, Stock: 20}
	updatedProductJSON, _ := json.Marshal(updatedProduct)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/api/products/:id", bytes.NewReader(updatedProductJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(product.ID)))

	if assert.NoError(t, handler.UpdateProduct(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var returnedProduct models.Product
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &returnedProduct)) {
			assert.Equal(t, updatedProduct.Name, returnedProduct.Name)
			assert.Equal(t, updatedProduct.Price, returnedProduct.Price)
			assert.Equal(t, updatedProduct.Stock, returnedProduct.Stock)
		}
	}
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := repositories.NewMockProductRepository()
	mockService := services.NewProductService(mockRepo)
	handler := handlers.NewProductHandler(mockService)

	product := &models.Product{Name: "Test Product", Price: 100, Stock: 10}
	mockRepo.Create(product)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/products/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(product.ID)))

	if assert.NoError(t, handler.DeleteProduct(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
