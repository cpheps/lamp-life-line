package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cpheps/lamp-life-line/cluster"
)

func colorHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		processColorPut(w, r)
	case http.MethodGet:
		processColorGet(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(formatErrorJson(fmt.Sprintf("Method %s not supported", r.Method)))
	}
}

func processColorPut(w http.ResponseWriter, r *http.Request) {
	jq, err := parseJSON(w, r)
	if err != nil {
		return
	}

	id, err := jq.String("id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatErrorJson(invalidRequest + ": property 'id' not found"))
		return
	}

	color, err := jq.Int("color")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatErrorJson(invalidRequest + ": property 'color' not found"))
		return
	}

	cluster, err := cluster.GetManagerInstance().GetCluster(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatErrorJson("Cluster not found"))
		return
	}

	color32 := uint32(color)
	cluster.Color = &color32

	bytes, err := json.Marshal(*cluster)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatErrorJson("getting Cluster"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func processColorGet(w http.ResponseWriter, r *http.Request) {
	jq, err := parseJSON(w, r)
	if err != nil {
		return
	}

	id, err := jq.String("id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatErrorJson(invalidRequest + ": property 'id' not found"))
		return
	}

	cluster, err := cluster.GetManagerInstance().GetCluster(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatErrorJson("Cluster not found"))
		return
	}

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("{\"color\": %d}", *cluster.Color))

	w.WriteHeader(http.StatusOK)
	w.Write(buffer.Bytes())
	return
}
