package routes

import (
	"go-echo-crud/handlers"
	"go-echo-crud/repositories"
	"go-echo-crud/services"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	productRepo := repositories.NewGormProductRepository()
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	e.POST("/products", productHandler.CreateProduct)
	e.GET("/products", productHandler.GetProducts)
	e.GET("/products/:id", productHandler.GetProduct)
	e.PUT("/products/:id", productHandler.UpdateProduct)
	e.DELETE("/products/:id", productHandler.DeleteProduct)
}
