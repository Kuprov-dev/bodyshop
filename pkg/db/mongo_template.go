package db

import (
	"bodyshop/pkg/conf"
	"bodyshop/pkg/models"
	"bytes"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTemplateDAO struct {
	db                  *mongo.Database
	templateCollection  string
	taskQueueCollection string
}

func NewMongoTemplateDAO(ctx context.Context, db *mongo.Database, config *conf.Config) *MongoTemplateDAO {
	return &MongoTemplateDAO{db: db, templateCollection: "templates"}
}

// This function aims to get data passed in headers,
// insert params to template and write it to db.
func (dao *MongoTemplateDAO) GetTemplateByUUID(ctx context.Context, templateUUID primitive.ObjectID) (*models.Template, error) {
	collection := dao.db.Collection(dao.templateCollection)

	var template models.Template

	filter := bson.M{"_id": templateUUID}

	err := collection.FindOne(ctx, filter).Decode(&template)
	if err != nil {
		return nil, err
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
