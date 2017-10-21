package server

import (
	"encoding/json"
	"github.com/cpheps/lamp-life-line/cluster"
	"io"
	"net/http"
)

func lampHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		processLampPost(w, r)
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
		w.WriteHeader(http.StatusBadRequest)
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
