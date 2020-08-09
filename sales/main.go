package main

import (
	"github.com/dtc/sales/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Restricted group
	r := e.Group("/api/v1")
	// Login route
	r.POST("/login", controllers.Login)
	r.POST("/products", controllers.CreateProduct)

	// Unauthenticated route
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/unaccess", controllers.AccessMe)

	e.Logger.Fatal(e.Start(":1323"))
}
