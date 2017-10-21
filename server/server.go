package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jmoiron/jsonq"
)

const (
	invalidRequest  = "Invalid JSON"
	clusterNotFound = "Cluster Not Found"
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
	http.HandleFunc("/cluster", clusterHandler)
	http.HandleFunc("/lamp", lampHandler)
	http.HandleFunc("/color", colorHandler)
}

func colorHandler(w http.ResponseWriter, r *http.Request) {
	//Get will return cluster color
	//Put will change color
}

func parseJSON(w http.ResponseWriter, r *http.Request) (*jsonq.JsonQuery, error) {
	resp := make(map[string]interface{})
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, invalidRequest)
		return nil, errors.New(invalidRequest)
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, invalidRequest)
		return nil, errors.New(invalidRequest)
	}

	jq := jsonq.NewQuery(resp)

	return jq, nil
}
