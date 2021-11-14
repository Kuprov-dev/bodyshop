package db

import (
	"bodyshop/pkg/conf"
	"bodyshop/pkg/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTasksQueueDAO struct {
	db                  *mongo.Database
	taskQueueCollection string
}

func (dao *MongoTasksQueueDAO) GetSendMailTaskByUUID(ctx context.Context, taskUUID primitive.ObjectID) (*models.SendMailTask, error) {
	collection := dao.db.Collection(dao.taskQueueCollection)

	var sendMailTask models.SendMailTask

	filter := bson.M{"_id": taskUUID}

	err := collection.FindOne(ctx, filter).Decode(&sendMailTask)
	if err != nil {
		return nil, err
	}

	return &sendMailTask, nil
}

func (dao *MongoTasksQueueDAO) CreateSendMailTask(ctx context.Context, taskData *models.SendMailTask) error {
	collection := dao.db.Collection(dao.taskQueueCollection)

	_, err := collection.InsertOne(ctx, taskData)
	if err != nil {
		return err
	}

	return nil
}

func NewMongoTasksQueueDAO(ctx context.Context, db *mongo.Database, config *conf.Config) *MongoTasksQueueDAO {
	return &MongoTasksQueueDAO{db: db, taskQueueCollection: config.TasksQueueDatabase.TasksQueueCollection}
}
