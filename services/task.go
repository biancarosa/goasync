package services

import (
	"net/http"
	"os"

	"github.com/biancarosa/goasync/models"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Setup Logrus
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

//TaskService is the interface that the describes all methods that a TaskService should have
type TaskService interface {
	ExecuteTask(task *models.Task)
	RetrieveTask(uuid uuid.UUID) *models.Task
}

type taskService struct{}

//NewTaskService creates a TaskService
func NewTaskService() TaskService {
	return new(taskService)
}

//ExecuteTask is the TaskService method that executes a task
func (s *taskService) ExecuteTask(task *models.Task) {
	log.Debug("Generating uuid")
	var err error
	task.UUID = uuid.NewV4()
	err = task.Create()
	if err != nil {
		log.WithFields(log.Fields{
			"task":  task,
			"error": err.Error(),
		}).Error("An error happened while creating the request.  The request has not been sent.")
		return
	}
	log.WithFields(log.Fields{
		"task": task,
	}).Debug("Execute task")
	go func() {
		task.UpdateStatus(task.Executing)
		client := new(http.Client)
		req, err := http.NewRequest(task.Method, task.URL, nil)
		if err != nil {
			log.WithFields(log.Fields{
				"task":  task,
				"error": err.Error(),
			}).Error("An error happened while creating the request.  The request has not been sent.")
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			log.WithFields(log.Fields{
				"task":  task,
				"error": err.Error(),
			}).Error("An error happened while executing the request. The request could have been sent.")
			task.UpdateStatus(task.Failure)
			return
		}
		log.WithFields(log.Fields{
			"task":     task,
			"response": resp.Body,
			"code":     resp.StatusCode,
		}).Info("Task execution finished.")
		task.UpdateStatus(task.Success)
	}()
}

//RetrieveTask is the TaskService method that retrieves a Task
func (s *taskService) RetrieveTask(uuid uuid.UUID) *models.Task {
	task := new(models.Task)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("Retrieve task")
	task.UUID = uuid
	err := task.Get()
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
