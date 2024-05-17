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

	e.POST("/api/products", productHandler.CreateProduct)
	e.GET("/api/products", productHandler.GetProducts)
	e.GET("/api/products/:id", productHandler.GetProduct)
	e.PUT("/api/products/:id", productHandler.UpdateProduct)
	e.DELETE("/api/products/:id", productHandler.DeleteProduct)
}
