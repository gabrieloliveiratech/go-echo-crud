package main

import (
	"fmt"
	"go-echo-crud/config"
	"go-echo-crud/models"
	"go-echo-crud/routes"
	"time"

	_ "go-echo-crud/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			Swagger Products CRUD API
//	@version		1.0
//	@description	This is a sample Products CRUD.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	https://github.com/gabrieloliveiratech
//	@contact.email	gabrielrbg8@gmail.com

//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// @host		localhost
// @BasePath	/api
func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	var attempts int
	for {
		err := connectDB()
		if err == nil {
			break
		}
		attempts++
		fmt.Printf("Failed to connect to database (attempt %d): %v\n", attempts, err)
		time.Sleep(2 * time.Second)
	}

	config.DB.AutoMigrate(&models.Product{})

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func connectDB() error {
	config.Connect()
	db, err := config.DB.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}
