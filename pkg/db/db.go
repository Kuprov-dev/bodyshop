package db

import (
	"bytes"
	"context"
	"encoding/json"
	"html/template"
	"io"

	"github.com/ddbelyaev/bodyshop/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TemplateDAO interface {
	/* TODO: add methods */
	GetTemplate(ctx context.Context, mClient *mongo.Client, body io.ReadCloser) (*models.Template, error)
	WriteTemplate(ctx context.Context, mClient *mongo.Client, t *models.Template) error
}

type MongoTemplateDAO struct {
}

// This function aims to get data passed in headers,
// insert params to template and write it to db.
func (*MongoTemplateDAO) GetTemplate(ctx context.Context, mClient *mongo.Client, body io.ReadCloser) (*models.Template, error) {
	mCollection := mClient.Database("mybd").Collection("template")

	var t models.Template

	err := json.NewDecoder(body).Decode(&t)
	if err != nil {
		return &t, err
	}

	filter := bson.D{
		{"template_uuid", bson.M{"$eq": t.TemplateUUID}},
	}

	var data models.Template
	err := mCollection.FindOne(ctx, filter).Decode(&t)
	if err != nil {
		panic(err)
	}

	// TODO: Where to get status? <31-10-21, ddbelyaev> //
	t.Status = "Pending"

	// TODO: Complete this method with 1. Getting template html 2. writing params to html <31-10-21, ddbelyaev> //
	var templatedMessage bytes.Buffer
	tmpl, _ := template.New("message template").Parse(data.Template)
	if err := tmpl.Execute(&templatedMessage, t.Params); err != nil {
		return &t, err
	}

	t.Template = templatedMessage.String()

	return &t, nil
}

// DONE: Change database and collection name <31-10-21, ddbelyaev> //
func (*MongoTemplateDAO) WriteTemplate(ctx context.Context, mClient *mongo.Client, t *models.Template) error {
	mCollection := mClient.Database("mydb").Collection("queue")
	_, err := mCollection.InsertOne(ctx, t)
	return err
}
