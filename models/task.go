package models

import (
	uuid "github.com/satori/go.uuid"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Task is the structure that contains all relevant information for a task
type Task struct {
	Method string
	URL    string
	Name   string
	UUID   uuid.UUID
}

//Create creates a task
func (task *Task) Create() error {
	session, err := mgo.Dial("db:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collection := session.DB("async").C("tasks")
	err = collection.Insert(task)
	return err
}

//Get returns a task based on its uuid
func (task *Task) Get(uuid string) error {
	session, err := mgo.Dial("db:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collection := session.DB("async").C("tasks")
	err = collection.Find(bson.M{"uuid": uuid}).One(&task)
	return err
}
