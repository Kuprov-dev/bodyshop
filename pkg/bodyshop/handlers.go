package bodyshop

import (
	"fmt"
	"net/http"

	"github.com/ddbelyaev/bodyshop/pkg/db"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateSendTask(mClient *mongo.Client, t *db.TemplateDAO) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if r.Body == nil {
				http.Error(w, "Please send a request body", 400)
				return
			}

			ctx := r.Context()

			template, err := (*t).GetTemplate(ctx, mClient, r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Println(err)
				return
			}

			err := (*t).WriteTemplate(ctx, mClient, template)
			if err != nil {
				fmt.Println(err)
				return
			}

		} else {
			fmt.Fprintf(w, "Only POST method supported for this route")
		}
	}
}
