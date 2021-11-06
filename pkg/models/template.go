package models

type Template struct {
	UUID         string   `bson:"_id,omitempty"`
	Status       string   `bson:"template"`
	TemplateUUID string   `bson:"template_uuid"`
	Params       []string `bson:"params"`
	Receivers    []string `bson:"receivers"`
	Message      string   `bson:"message"`
	Template     string   `bson:"template"`
}
