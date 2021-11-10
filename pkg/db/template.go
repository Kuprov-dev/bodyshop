package db

import (
	"bodyshop/pkg/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TemplateDAO interface {
	GetTemplateByUUID(ctx context.Context, templateUUID primitive.ObjectID) (*models.Template, error)
}
