package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cpheps/lamp-life-line/cluster"
)

func clusterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		processClusterPost(w, r)
	case http.MethodGet:
		processClusterGet(w, r)
	case http.MethodDelete:
		processClusterDelete(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(formatErrorJson(fmt.Sprintf("Method %s not supported", r.Method)))
	}
}

func processClusterPost(w http.ResponseWriter, r *http.Request) {
	jq, err := parseJSON(w, r)
	if err != nil {
		return
	}

	name, err := jq.String("name")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatErrorJson(invalidRequest + ": property 'name' not found"))
		return
	}

	color, err := jq.Int("color")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatErrorJson(invalidRequest + ": property 'color' not found"))
		return
	}
	color32 := int32(color)

	cluster := cluster.GetManagerInstance().RegisterNewCluster(name, color32)

	bytes, err := json.Marshal(*cluster)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatErrorJson("Error registering new Cluster"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func processClusterGet(w http.ResponseWriter, r *http.Request) {
	jq, err := parseJSON(w, r)
	if err != nil {
		return
	}

	id, err := jq.String("id")
	if err != nil {
		clusters := cluster.GetManagerInstance().GetClusters()
		bytes, err := json.Marshal(clusters)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatErrorJson("Error getting Clusters"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
		return
	}

	cluster, err := cluster.GetManagerInstance().GetCluster(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(formatErrorJson("Error Cluster not found"))
		return
	}

	bytes, err := json.Marshal(*cluster)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatErrorJson("Error getting Cluster"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func processClusterDelete(w http.ResponseWriter, r *http.Request) {
	jq, err := parseJSON(w, r)
	if err != nil {
		return
	}

	id, err := jq.String("id")
	if err != nil {
		clusters := cluster.GetManagerInstance().GetClusters()
		bytes, err := json.Marshal(clusters)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatErrorJson("Error getting Clusters"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
		return
	}

	_, err = cluster.GetManagerInstance().UnregisterCluster(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(formatErrorJson(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(formatErrorJson("Successfully unregistred cluster"))
}
