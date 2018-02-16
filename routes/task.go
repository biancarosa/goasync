package routes

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

//TaskRequest is the struct that defines a request
type TaskRequest struct {
	Method string
	URL    string
}

//Task is the route that executes a request to the URL of the task with the name same as a the task URL argument
func Task(ctx echo.Context) error {
	task := ctx.QueryParam("task")
	req := new(TaskRequest)
	log.Debug("Binding task request")
	if err := ctx.Bind(req); err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"task": task,
		"req":  req,
	}).Debug("Execute task")
	go func() {
		client := new(http.Client)
		req, err := http.NewRequest("GET", "http://example.com", nil)
		if err != nil {
			log.WithFields(log.Fields{
				"task":  task,
				"req":   req,
				"error": err.Error(),
			}).Error("An error happened while creating the request.")
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			log.WithFields(log.Fields{
				"task":  task,
				"req":   req,
				"error": err.Error(),
			}).Error("An error happened while executing the request.")
			return
		}
		log.WithFields(log.Fields{
			"task":     task,
			"req":      req,
			"response": resp.Body,
		}).Info("Everything went fine with the task.")
	}()
	return ctx.String(http.StatusOK, "Your task is being executed.")
}
