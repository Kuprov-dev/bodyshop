package main

import (
	"bodyshop/pkg/bodyshop"
	logging "bodyshop/pkg/log"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {
	app := bodyshop.NewApp(context.Background())

	r := mux.NewRouter()
	r.Handle("/create_send_task", app.CreateSendTask()).Methods(http.MethodPost)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	handler := logging.LoggingMiddleware(app.Logger)(r)

	s := &http.Server{
		Addr:    app.Config.ServerAddr(),
		Handler: handler,
	}
	defer s.Close()

	go func() {
		log.Println("Server started...")

		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)

			return
		}
	}()

	<-stop

	log.Println("Server stopped...")
}
