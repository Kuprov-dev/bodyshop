package db

import (
	"bodyshop/pkg/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskQueueDAO interface {
	GetSendMailTaskByUUID(ctx context.Context, taskUUID primitive.ObjectID) (*models.SendMailTask, error)
	CreateSendMailTask(ctx context.Context, taskData *models.SendMailTask) error
}
