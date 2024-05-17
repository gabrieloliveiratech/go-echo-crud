package handlers

import (
	"net/http"
	"strconv"

	"go-echo-crud/models"
	"go-echo-crud/responses"
	"go-echo-crud/services"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// @Summary		Create a new product
// @Description	Create a new product with the given body
// @ID				create-product
// @Tags			Products
// @Accept			json
// @Produce		json
// @Param			product	body		models.Product	true	"Product body"
// @Success		201		{object}	models.Product
// @Router			/products [post]
func (h *ProductHandler) CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return responses.Error(c, http.StatusBadRequest, err)
	}
	if err := h.service.CreateProduct(product); err != nil {
		return responses.Error(c, http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, product)
}

// @Summary		Get all products
// @Description	Get a list of all products
// @ID				get-products
// @Tags			Products
// @Produce		json
// @Success		200	{array}	models.Product
// @Router			/products [get]
func (h *ProductHandler) GetProducts(c echo.Context) error {
	products, err := h.service.GetProducts()
	if err != nil {
		return responses.Error(c, http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, products)
}

// @Summary		Get a product
// @Description	Get a product with the given ID
// @ID				get-product
// @Tags			Products
// @Produce		json
// @Param			id	path		int	true	"Product ID"
// @Success		200	{object}	models.Product
// @Router			/products/{id} [get]
func (h *ProductHandler) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.service.GetProduct(uint(id))
	if err != nil {
		return responses.NotFound(c, "Product not found")
	}
	return c.JSON(http.StatusOK, product)
}

// @Summary		Update a product
// @Description	Update a product with the given ID and body
// @ID				update-product
// @Tags			Products
// @Accept			json
// @Produce		json
// @Param			id		path		int						true	"Product ID"
// @Param			product	body		models.ProductUpdate	true	"Product body"
// @Success		200		{object}	models.Product
// @Router			/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	updatedProduct := &models.ProductUpdate{}
	if err := c.Bind(updatedProduct); err != nil {
		return responses.Error(c, http.StatusBadRequest, err)
	}

	product, err := h.service.UpdateProduct(uint(id), updatedProduct)
	if err != nil {
		return responses.Error(c, http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, product)
}

// @Summary		Delete a product
// @Description	Delete a product with the given ID
// @ID				delete-product
// @Tags			Products
// @Produce		json
// @Param			id	path		int	true	"Product ID"
// @Success		204	{object}	models.Product
// @Router			/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteProduct(uint(id)); err != nil {
		return responses.Error(c, http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
