package routes

import (
	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type model struct {
	Method string
	URL    string
	Name   string
	UUID   uuid.UUID
}

//Task is the route that executes a request to the URL of the task with the name same as a the task URL argument
func Task(ctx echo.Context) error {
	task := new(model)
	log.Debug("Binding task request")
	if err := ctx.Bind(task); err != nil {
		return err
	}
	log.Debug("Generating uuid")
	task.UUID = uuid.Must(uuid.NewV4())
	log.WithFields(log.Fields{
		"task": task,
	}).Debug("Execute task")
	go func() {
		client := new(http.Client)
		req, err := http.NewRequest(task.Method, task.URL, nil)
		if err != nil {
			log.WithFields(log.Fields{
				"task":  task,
				"error": err.Error(),
			}).Error("An error happened while creating the request.")
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			log.WithFields(log.Fields{
				"task":  task,
				"error": err.Error(),
			}).Error("An error happened while executing the request.")
			return
		}
		log.WithFields(log.Fields{
			"task":     task,
			"response": resp.Body,
		}).Info("Everything went fine with the task.")
	}()
	return ctx.JSON(http.StatusCreated, task)
}
