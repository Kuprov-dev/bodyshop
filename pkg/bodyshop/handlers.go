package bodyshop

import (
	"bodyshop/pkg/schemes"
	"encoding/json"
	"log"
	"net/http"
)

func (app *App) CreateSendTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}

		var createTaskRequest schemes.CreateSendTaskRequest
		if err := json.NewDecoder(r.Body).Decode(&createTaskRequest); err != nil {
			http.Error(w, "Error decoding request", 400)

			return
		}

		_, err := CreateSendMailTask(r.Context(), createTaskRequest, app.TemplateDAO, app.TasksQueueDAO)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		w.Write([]byte("Ok"))
	}
}
