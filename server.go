package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/biancarosa/goasync/routes"
)

func main() {
	e := echo.New()

	//Setup Logrus
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", routes.HealthCheck)
	e.POST("/async", routes.Task)

	e.Logger.Fatal(e.Start(":1323"))
}
