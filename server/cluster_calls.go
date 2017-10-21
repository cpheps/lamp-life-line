package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cpheps/lamp-life-line/cluster"
)

func clusterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		processClusterPost(w, r)
	case http.MethodGet:
		processClusterGet(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, fmt.Sprintf("Method %s not supported", r.Method))
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
		io.WriteString(w, invalidRequest+": property 'name' not found")
		return
	}

	cluster := cluster.GetManagerInstance().RegisterNewCluster(name)

	bytes, err := json.Marshal(*cluster)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error registering new Cluster")
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
			io.WriteString(w, "Error getting Clusters")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
		return
	}

	cluster, err := cluster.GetManagerInstance().GetCluster(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Error Cluster not found")
		return
	}

	bytes, err := json.Marshal(*cluster)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error getting Cluster")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
