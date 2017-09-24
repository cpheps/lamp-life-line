package server

import (
	"fmt"
	"net/http"
	"os"
)

func StartServer() <-chan error {
	fmt.Println("Starting Server")
	errors := make(chan error, 1)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	initHandlers()

	go func() {
		defer close(errors)
		errors <- http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	}()

	return errors
}

func initHandlers() {
	http.HandleFunc("/register-cluster", registerClusterHandler)
	http.HandleFunc("/register-lamp", registerLampHandlers)
	http.HandleFunc("/color", colorHandler)
}

//Handlers
func registerClusterHandler(w http.ResponseWriter, r *http.Request) {
	//Post will create cluster and return json for it
}

func registerLampHandlers(w http.ResponseWriter, r *http.Request) {
	//Post will create lamp and return json for it
}

func colorHandler(w http.ResponseWriter, r *http.Request) {
	//Get will return cluster color
	//Put will change color
}
