package routes

import (
	"github.com/biancarosa/goasync/config"
	"net/http"

	"github.com/labstack/echo"
)

//Task is the route that executes a request to the URL of the task with the name same as a the task URL argument
func Task(ctx echo.Context) error {
	config := config.ParseTOML()
	task := ctx.QueryParam("task")
	for _, endpoint := range config.Endpoints {
		if endpoint.Name == task {
			print(endpoint.Name)
		}
	}
	return ctx.String(http.StatusOK, "Your task is being executed.")
}
