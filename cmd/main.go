package main

import (
	"bodyshop/pkg/bodyshop"
	"bodyshop/pkg/conf"
	"bodyshop/pkg/db"
	logging "bodyshop/pkg/log"
	"bodyshop/pkg/providers"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	PORT             string = ":8080"
	connectionString string = "mongodb://localhost:27017/"
)

func main() {
	config := conf.New()
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	logEntry := logrus.NewEntry(log)

	mailsenderService := &providers.HTTPMailsenderServiceProvider{Config: config}
	templateDAO := &db.MongoTemplateDAO{}

	app := bodyshop.App{
		Config:            config,
		MailsenderService: mailsenderService,
		Logger:            logEntry,
		TemplateDAO:       templateDAO,
	}

	r := mux.NewRouter()
	r.Handle("/create_send_task", app.CreateSendTask()).Methods(http.MethodPost)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	handler := logging.LoggingMiddleware(logEntry)(r)

	s := &http.Server{
		Addr:    PORT,
		Handler: handler,
	}
	defer s.Close()

	go func() {
		fmt.Println("Server started...")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
			return
		}
	}()

	<-stop

	fmt.Println("Server stopped...")
}
