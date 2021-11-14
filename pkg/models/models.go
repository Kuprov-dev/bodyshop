package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HTMLTeplate struct {
	UUID     primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Template string             `json:"template"`
	Params   []string           `json:"params"`
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
	ParamsMap    map[string]string  `bson:"params_map"`
	Receivers    []string           `bson:"receivers"`
	Message      string             `bson:"message"`
	Username     string             `bson:"username"`
}
