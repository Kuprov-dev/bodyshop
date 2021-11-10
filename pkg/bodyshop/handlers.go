package bodyshop

import (
	"bodyshop/pkg/schemes"
	"encoding/json"
	"fmt"
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

		err := createSendMailTask(r.Context(), createTaskRequest)

		template, err := app.TemplateDAO.GetTemplate(r.Context(), templateUUID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		err = t.WriteTemplate(ctx, mClient, template)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
