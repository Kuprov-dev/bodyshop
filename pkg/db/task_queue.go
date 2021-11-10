package db

import (
	"bodyshop/pkg/models"
	"bodyshop/pkg/schemes"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskQueueDAO interface {
	GetSendMailTaskByUUID(ctx context.Context, taskUUID primitive.ObjectID) (*models.Template, error)
	CreateSendMailTask(ctx context.Context, taskData *schemes.QueueCreateSendTaskRequest) error
}
