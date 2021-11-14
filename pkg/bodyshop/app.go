package bodyshop

import (
	"bodyshop/pkg/conf"
	"bodyshop/pkg/db"
	"context"
	"log"

	"github.com/sirupsen/logrus"
)

type App struct {
	Config        *conf.Config
	Logger        *logrus.Entry
	TemplateDAO   db.TemplateDAO
	TasksQueueDAO db.TaskQueueDAO
}

func NewApp(ctx context.Context) *App {
	log.Println("App init...")
	config := conf.New()

	templatesDBConnection, err := db.GetTemplatesDBConnection(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	tasksQueueDBConnection, err := db.GetTasksQueueDBConnection(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	templateDAO := db.NewMongoTemplateDAO(ctx, templatesDBConnection, config)
	tasksQueueDAO := db.NewMongoTasksQueueDAO(ctx, tasksQueueDBConnection, config)

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	logEntry := logrus.NewEntry(log)

	defer log.Println("App init done.")

	return &App{
		Config:        config,
		Logger:        logEntry,
		TemplateDAO:   templateDAO,
		TasksQueueDAO: tasksQueueDAO,
	}
}
