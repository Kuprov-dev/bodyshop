package db

import (
	"bodyshop/pkg/conf"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	templatesDBConnection  *mongo.Database
	tasksQueueDBConnection *mongo.Database
)

func ConnectMongoDB(ctx context.Context, dbURI string, dbName string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	dbConnection := client.Database(dbName)

	log.Println("Connection to DB success")

	return dbConnection, nil
}

func GetTemplatesDBConnection(ctx context.Context, config *conf.Config) (*mongo.Database, error) {
	var err error

	if templatesDBConnection == nil {
		templatesDBConnection, err = ConnectMongoDB(ctx, config.TemplatesDatabaseURI(), config.TemplatesDBName())
	}

	return templatesDBConnection, err
}

func GetTasksQueueDBConnection(ctx context.Context, config *conf.Config) (*mongo.Database, error) {
	var err error

	if tasksQueueDBConnection == nil {
		tasksQueueDBConnection, err = ConnectMongoDB(ctx, config.TasksQueueCollection(), config.TasksQueueDBName())
	}

	return tasksQueueDBConnection, err
}
