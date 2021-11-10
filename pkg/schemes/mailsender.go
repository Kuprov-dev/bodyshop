package schemes

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateSendTaskRequest struct {
	TemplateUUID primitive.ObjectID `json:"template_uuid"`
	Params       map[string]string  `json:"params"`
	Username     string             `json:"username"`
	Receivers    []string           `json:"receivers"`
}

type QueueCreateSendTaskRequest struct {
	TemplateUUID primitive.ObjectID `json:"template_uuid"`
	Params       map[string]string  `json:"params"`
	Username     string             `json:"username"`
	Receivers    []string           `json:"receivers"`
}
