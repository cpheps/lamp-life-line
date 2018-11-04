package common

import (
	"testing"
)

func TestGenerateFailurePayload(t *testing.T) {
	expected := `{"error":"my error"}`

	errorMessage := "my error"

	data, err := generateFailurePayload(errorMessage)
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	json := string(data)

	if json != expected {
		t.Errorf("Expected\n\n%s\n\ngot\n\n%s", expected, json)
	}
}
