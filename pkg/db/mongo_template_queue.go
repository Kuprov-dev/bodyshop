package db

import (
	"bodyshop/pkg/conf"
	"bodyshop/pkg/models"
	"bodyshop/pkg/schemes"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoTaskQueueDAO struct {
	config *conf.Config
}

func (dao *MongoTaskQueueDAO) GetSendMailTaskByUUID(ctx context.Context, taskUUID primitive.ObjectID) (*models.Template, error) {
	return nil, nil
}

func (dao *MongoTaskQueueDAO) CreateSendMailTask(ctx context.Context, taskData *schemes.QueueCreateSendTaskRequest) error {
	return nil
}
