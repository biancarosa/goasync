package models

import (
	uuid "github.com/satori/go.uuid"

	"gopkg.in/mgo.v2"
)

type Task struct {
	Method string
	URL    string
	Name   string
	UUID   uuid.UUID
}

func (task *Task) Create() error {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("test").C("people")
	err = collection.Insert(task)
	return err
}
