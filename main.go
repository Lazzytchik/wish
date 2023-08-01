package main

import (
	"net/http"
	"os"

	"log"

	"github.com/lazzytchik/wish/app"
)

func main() {

	l := log.New(os.Stdout, "[SERVER]: ", log.Lmicroseconds)

	srv := app.Service{
		Logger: l,
		Api:    &app.Usecases{},
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: srv.Routes(),
	}

	srv.Server = server

	server.ListenAndServe()
}
