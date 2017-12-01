package server

import (
	"encoding/json"
	"fmt"
	"github.com/cpheps/lamp-life-line/cluster"
	"io"
	"net/http"
)

func lampHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		processLampPost(w, r)
	case http.MethodPut:
		processLampPut(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, fmt.Sprintf("Method %s not supported", r.Method))
	}
}

func processLampPost(w http.ResponseWriter, r *http.Request) {
	jq, err := parseJSON(w, r)
	if err != nil {
		return
	}

	id, err := jq.String("id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, invalidRequest+": property 'id' not found")
		return
	}

	clusterID, err := jq.String("clusterId")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, invalidRequest+": property 'clusterId' not found")
		return
	}

	cluster, err := cluster.GetManagerInstance().GetCluster(clusterID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Cluster Not Found")
		return
	}

	lamp, err := cluster.RegisterNewLamp(id, r.RemoteAddr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	bytes, err := json.Marshal(lamp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error registering Lamp")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func processLampPut(w http.ResponseWriter, r *http.Request) {
	jq, err := parseJSON(w, r)
	if err != nil {
		return
	}

	id, err := jq.String("id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, invalidRequest+": property 'id' not found")
		return
	}

	clusterID, err := jq.String("clusterId")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, invalidRequest+": property 'clusterId' not found")
		return
	}

	cluster, err := cluster.GetManagerInstance().GetCluster(clusterID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Cluster Not Found")
		return
	}

	lamp, err := cluster.GetLamp(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Lamp Not Found in Cluster"+*cluster.ID)
		return
	}

	lamp.ListenAddress = &r.RemoteAddr
	//modify color
	w.WriteHeader(http.StatusOK)
}
