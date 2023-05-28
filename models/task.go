package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
	Completed   bool               `json:"completed,omitempty"`
}

func (task Task) IsEmpty() bool {
	return task.Title == "" || task.Description == ""
}
