package routes

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/biancarosa/goasync/models"
	"github.com/biancarosa/goasync/services"
)

var service services.TaskService

func init() {
	service = services.NewTaskService()
}

//Task is the route that executes a request to the URL of the task with the name same as a the task URL argument
func Task(ctx echo.Context) error {
	task := new(models.Task)
	log.Debug("Binding task request")
	if err := ctx.Bind(task); err != nil {
		return err
	}
	service.ExecuteTask(task)
	return ctx.JSON(http.StatusCreated, task)
}
