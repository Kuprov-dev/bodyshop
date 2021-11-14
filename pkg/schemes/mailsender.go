package schemes

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateSendTaskRequest struct {
	TemplateUUID primitive.ObjectID `json:"template_uuid"`
	ParamsMap    map[string]string  `json:"params_map"`
	Username     string             `json:"username"`
	Receivers    []string           `json:"receivers"`
}
