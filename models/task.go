package models

import (
	uuid "github.com/satori/go.uuid"
)

type Task struct {
	Method string
	URL    string
	Name   string
	UUID   uuid.UUID
}
