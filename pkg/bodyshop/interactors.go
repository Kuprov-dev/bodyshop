package bodyshop

import (
	"bodyshop/pkg/db"
	"context"
	"fmt"
	"net/http"
)

// Интерактор, для создания таски на отправку мейла в очередь
func createSendMailTask(ctx context.Context, username string, templateDAO db.TemplateDAO) (*schemes.UserRecievers, error) {
	template, err := templateDAO.GetTemplate(r.Context(), templateUUID)
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
