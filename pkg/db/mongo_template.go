package db

import (
	"bodyshop/pkg/conf"
	"bodyshop/pkg/models"
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

func (dao *MongoTemplateDAO) GetTemplateByUUID(ctx context.Context, templateUUID primitive.ObjectID) (*models.HTMLTeplate, error) {
	collection := dao.db.Collection(dao.templateCollection)

	var template models.HTMLTeplate

	filter := bson.M{"_id": templateUUID}

	err := collection.FindOne(ctx, filter).Decode(&template)
	if err != nil {
		return nil, err
	}

	return &template, nil
}
