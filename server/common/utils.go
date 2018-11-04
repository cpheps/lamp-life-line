// Package common contains common functions across all api version
package common

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// ReadBody read and parse the request body for a command job
func ReadBody(body io.ReadCloser, v interface{}) error {
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("Failed to read request body: %s", err.Error())
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		log.Printf("Failed to parse request body: %s", err.Error())
		return err
	}

	return nil
}

// WriteError write an error with the given status code and message back to a response
func WriteError(statusCode int, errorMessage string, w http.ResponseWriter) {
	data, err := generateFailurePayload(errorMessage)
	if err != nil {
		log.Printf("Failed to generate failure payload: %s", err.Error())
		return
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(data)
	log.Printf("Sent %d failure", statusCode)

	if err != nil {
		log.Printf("Failed to write error response with message '%s': %s", errorMessage, err.Error())
	}
}

// WriteSuccess writes the successful json payload with a 200 message
func WriteSuccess(data []byte, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(data)

	if err != nil {
		log.Printf("Failed to write ok response with message '%s': %s", string(data), err.Error())
	}
}
