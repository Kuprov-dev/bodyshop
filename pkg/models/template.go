package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Template struct {
	UUID         primitive.ObjectID `bson:"_id,omitempty"`
	Status       string             `bson:"template"`
	TemplateUUID primitive.ObjectID `bson:"template_uuid"`
	Params       []string           `bson:"params"`
	Receivers    []string           `bson:"receivers"`
	Message      string             `bson:"message"`
	Template     string             `bson:"template"`
}

const (
	Done = iota
	Pending
	Error
	Incoming
)

type SendMailTask struct {
	UUID         primitive.ObjectID `bson:"_id,omitempty"`
	Status       int                `bson:"status"`
	TemplateUUID primitive.ObjectID `bson:"template_uuid"`
	Params       []string           `bson:"params"`
	Receivers    []string           `bson:"receivers"`
	Message      string             `bson:"message"`
	Username     string             `bson:"username"`
}
