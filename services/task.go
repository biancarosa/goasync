package services

import (
	"fmt"
	"net/http"
	"os"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"

	"github.com/biancarosa/goasync/configuration"
	"github.com/biancarosa/goasync/models"
)

var conf *configuration.Configuration

func init() {
	// Setup Logrus
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	//Load configurations
	loader := new(configuration.EnvironmentLoader)
	var err error
	conf, err = loader.LoadConfiguration()
	if err != nil {
		panic("Could not load configurations")
	}
}

//TaskService is the interface that the describes all methods that a TaskService should have
type TaskService interface {
	ExecuteTask(task *models.Task)
	RetrieveTask(uuid uuid.UUID) *models.Task
}

type taskService struct {
	session *mgo.Session
}

//NewTaskService creates a TaskService
func NewTaskService() TaskService {
	ts := new(taskService)
	ts.session = newDatabaseSession()
	return ts
}

func newDatabaseSession() *mgo.Session {
	url := fmt.Sprintf("%s:%s", conf.MongoDB.Host, conf.MongoDB.Port)
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	return session
}

//ExecuteTask is the TaskService method that executes a task
func (s *taskService) ExecuteTask(task *models.Task) {
	log.Debug("Generating uuid")
	var err error
	task.UUID = uuid.NewV4()
	err = task.Create(s.session)
	defer s.session.Close()
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
			return
		}
		log.WithFields(log.Fields{
			"task":     task,
			"response": resp.Body,
			"code":     resp.StatusCode,
		}).Info("Task execution finished.")
	}()
}

//RetrieveTask is the TaskService method that retrieves a Task
func (s *taskService) RetrieveTask(uuid uuid.UUID) *models.Task {
	task := new(models.Task)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("Retrieve task")
	err := task.Get(s.session, uuid)
	defer s.session.Close()
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
