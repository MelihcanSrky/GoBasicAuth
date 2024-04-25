package main

import (
	"log"
	"net/http"

	"github.com/MelihcanSrky/GoBasicAuth/pkg/config"
	"github.com/MelihcanSrky/GoBasicAuth/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	app := new(config.AppConfig)
	app.Auth.Username = "Admin"
	app.Auth.Password = "123123"

	mux := chi.NewRouter()

	mux.HandleFunc("/home", BasicAuth(handlers.Home, app))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	log.Fatal(err)
}
