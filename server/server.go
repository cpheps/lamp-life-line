package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jmoiron/jsonq"
)

const (
	invalidRequest  = "Invalid JSON"
	clusterNotFound = "Cluster Not Found"
)

//StartServer creats a new server on the given port
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
	http.HandleFunc("/color", colorHandler)
}

func parseJSON(w http.ResponseWriter, r *http.Request) (*jsonq.JsonQuery, error) {
	resp := make(map[string]interface{})
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatErrorJson(invalidRequest))
		return nil, errors.New(invalidRequest)
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatErrorJson(invalidRequest))
		return nil, errors.New(invalidRequest)
	}

	jq := jsonq.NewQuery(resp)

	return jq, nil
}

func formatErrorJson(message string) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("{\"error\":\"%s\"}", message))
	return buffer.Bytes()
}
