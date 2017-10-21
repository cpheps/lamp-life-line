package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/cpheps/lamp-life-line/cluster"
	"github.com/jmoiron/jsonq"
)

func clusterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		processClusterPost(w, r)
	case http.MethodGet:
		processClusterGet(w, r)
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
	}

	cluster := cluster.GetManagerInstance().RegisterNewCluster(name)

	bytes, err := json.Marshal(*cluster)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error registering new Cluster")
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
		}
	}

	cluster, err := cluster.GetManagerInstance().GetCluster(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Error Cluster not found")
	}

	bytes, err := json.Marshal(*cluster)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error getting Cluster")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
