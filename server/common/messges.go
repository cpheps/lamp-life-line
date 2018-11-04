package common

import (
	"encoding/json"
)

// ColorMessage represents the color payload from a color request
type ColorMessage struct {
	Color uint32 `json:"color"`
}

// failureMessage represents the failure payload sent back when a request fails
type failureMessage struct {
	Error string `json:"error"`
}

func generateFailurePayload(message string) (data []byte, err error) {
	failure := &failureMessage{
		Error: message,
	}

	data, err = json.Marshal(failure)
	return
}
