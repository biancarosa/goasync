package routes

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Setup Logrus
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

//HealthCheck is the route that prints a sucessful message when the application is fine.
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "I seem to be perfectly fine.")
}
