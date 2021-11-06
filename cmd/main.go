package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ddbelyaev/bodyshop/pkg/bodyshop"
	"github.com/ddbelyaev/bodyshop/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PORT string = ":8080"
const connectionString string = "mongodb://localhost:27017/"

func main() {
	ctx := context.Background()

	cOpts := options.Client().ApplyURI(connectionString)
	mClient, err := mongo.Connect(ctx, cOpts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	t := db.MongoTemplateDAO{}

	mux := http.NewServeMux()

	fmt.Println("Server started...")

	mux.Handle("/create_send_task", bodyshop.CreateSendTask(&mClient, &t))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	s := &http.Server{
		Addr:    PORT,
		Handler: mux,
	}
	defer s.Close()

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
			return
		}
	}()

	<-stop

	fmt.Println("Server stopped...")

}
