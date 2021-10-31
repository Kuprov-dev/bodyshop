package db

import (
	"context"
	"encoding/json"
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

func (*MongoTemplateDAO) GetTemplate(ctx context.Context, mClient *mongo.Client, body io.ReadCloser) (*models.Template, error) {
	mCollection := mClient.Database("library").Collection("books")

	var template models.Template

	err := json.NewDecoder(body).Decode(&template)
	if err != nil {
		return &template, err
	}

	filter := bson.D{
		{"template_uuid", bson.M{"$eq": template.TemplateUUID}},
	}

	var data models.Template
	err := mCollection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		panic(err)
	}

	// TODO: Where to get status? <31-10-21, ddbelyaev> //
	template.Status = "Pending"

	// TODO: Complete this method with 1. Getting template html 2. writing params to html <31-10-21, ddbelyaev> //

	return &template, nil
}

// TODO: Change database and collection name <31-10-21, ddbelyaev> //
func (*MongoTemplateDAO) WriteTemplate(ctx context.Context, mClient *mongo.Client, t *models.Template) error {
	mCollection := mClient.Database("library").Collection("books")
	_, err := mCollection.InsertOne(ctx, t)
	return err

}
