package services

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	"github.com/biancarosa/goasync/models"
)

type TaskService interface {
	ExecuteTask(task *models.Task)
	RetrieveTask(uuid string) *models.Task
}
type taskService struct{}

func NewTaskService() TaskService {
	return new(taskService)
}

func (s *taskService) ExecuteTask(task *models.Task) {
	log.Debug("Generating uuid")
	task.UUID = uuid.Must(uuid.NewV4())
	task.Create()
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
}

func (s *taskService) RetrieveTask(uuid string) *models.Task {
	task := new(models.Task)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("Retrieve task")
	err := task.Get(uuid)
	if err != nil {
		log.WithFields(log.Fields{
			"task":  task,
			"error": err.Error(),
		}).Error("An error happened while executing the query.")
		return nil
	}
	log.WithFields(log.Fields{
		"task": task,
	}).Info("Everything went fine finding the task.")
	return task
}
