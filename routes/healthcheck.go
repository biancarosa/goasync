package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

//HealthCheck is the route that prints a sucessful message when the application is fine.
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "I seem to be perfectly fine.")
}
