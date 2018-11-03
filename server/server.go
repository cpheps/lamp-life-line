package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cpheps/lamp-life-line/server/v1"
	"github.com/gorilla/mux"
)

//StartServer creats a new server on the given port
func StartServer() <-chan error {
	log.Printf("Starting Server")
	errors := make(chan error, 1)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	go func() {
		defer close(errors)
		errors <- http.ListenAndServe(fmt.Sprintf(":%s", port), initRouters())
	}()

	return errors
}

func initRouters() http.Handler {
	r := mux.NewRouter()

	//v1 routes
	v1.AddRoutes(r.PathPrefix("/v1").Subrouter())

	return r
}
