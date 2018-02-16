package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/biancarosa/goasync/routes"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", routes.HealthCheck)
	e.GET("/async/:task", routes.Task)

	e.Logger.Fatal(e.Start(":1323"))
}
