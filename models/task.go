package models

import (
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/biancarosa/goasync/configuration"
)

var conf *configuration.Configuration

func init() {
	// Setup Logrus
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	conf := new(configuration.Configuration)
}

//Task is the structure that contains all relevant information for a task
type Task struct {
	Method string
	URL    string
	Name   string
	UUID   uuid.UUID
}

//Create creates a task
func (task *Task) Create() (err error) {
	session, err := mgo.Dial(fmt.Sprintf("%s:%s", conf.MongoDB.Host, conf.MongoDB.Port))
	if err != nil {
		return
	}
	defer session.Close()

	collection := session.DB("async").C("tasks")
	return collection.Insert(task)
}

//Get returns a task based on its uuid
func (task *Task) Get(uuid uuid.UUID) (err error) {
	session, err := mgo.Dial("db:27017")
	if err != nil {
		return
	}
	defer session.Close()

	collection := session.DB("async").C("tasks")
	return collection.Find(bson.M{"uuid": uuid}).One(&task)
}
