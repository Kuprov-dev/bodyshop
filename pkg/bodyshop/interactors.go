package bodyshop

import (
	"bodyshop/pkg/db"
	"bodyshop/pkg/models"
	"bodyshop/pkg/schemes"
	"bytes"
	"context"
	"html/template"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func pushTaskInQueue(
	ctx context.Context,
	username string,
	paramsMap map[string]string,
	receivers []string,
	templatedMessage string,
	templateUUID primitive.ObjectID,
	tasksQueueDAO db.TaskQueueDAO,
) (primitive.ObjectID, error) {
	taskUUID := primitive.NewObjectID()
	sendMailTask := models.SendMailTask{
		UUID:         taskUUID,
		Status:       models.Incoming,
		TemplateUUID: templateUUID,
		ParamsMap:    paramsMap,
		Receivers:    receivers,
	}

	err := tasksQueueDAO.CreateSendMailTask(ctx, &sendMailTask)
	if err != nil {
		return taskUUID, err
	}

	return taskUUID, nil
}

func formatTemplatedMessage(
	htmlTemplate *models.HTMLTeplate,
	params map[string]string,
) (string, error) {
	var templatedMessage bytes.Buffer

	tmpl, err := template.New(htmlTemplate.Name).Parse(htmlTemplate.Template)
	if err != nil {
		return "", err
	}

	if err := tmpl.Execute(&templatedMessage, params); err != nil {
		return "", err
	}

	return templatedMessage.String(), nil
}

// Интерактор, для создания таски на отправку мейла в очередь
func CreateSendMailTask(
	ctx context.Context,
	createTaskRequest schemes.CreateSendTaskRequest,
	templateDAO db.TemplateDAO,
	tasksQueueDAO db.TaskQueueDAO,
) (primitive.ObjectID, error) {
	var taskUUID primitive.ObjectID

	htmlTemplate, err := templateDAO.GetTemplateByUUID(ctx, createTaskRequest.TemplateUUID)
	if err != nil {
		return taskUUID, err
	}

	templatedMessage, err := formatTemplatedMessage(htmlTemplate, createTaskRequest.ParamsMap)
	if err != nil {
		return taskUUID, err
	}

	taskUUID, err = pushTaskInQueue(
		ctx,
		createTaskRequest.Username,
		createTaskRequest.ParamsMap,
		createTaskRequest.Receivers,
		templatedMessage,
		htmlTemplate.UUID,
		tasksQueueDAO,
	)
	if err != nil {
		return taskUUID, err
	}

	return taskUUID, nil
}
