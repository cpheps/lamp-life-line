//Package v1 handles version 1 API endpoints
package v1

import (
	"encoding/json"
	"net/http"

	"github.com/cpheps/lamp-life-line/cluster"
	"github.com/cpheps/lamp-life-line/server/common"
	"github.com/gorilla/mux"
)

//AddRoutes adds all handlers to router
func AddRoutes(r *mux.Router) {
	r.HandleFunc("/cluster", handleClusterPost).
		Methods(http.MethodPost).
		HeadersRegexp("Content-Type", "application/json")

	r.HandleFunc("/cluster/{name:[a-zA-Z0-9_-]+}", handleClusterGet).
		Methods(http.MethodGet)

	r.HandleFunc("/cluster/{name:[a-zA-Z0-9_-]+}/color", handleColorGet).
		Methods(http.MethodGet)

	r.HandleFunc("/cluster/{name:[a-zA-Z0-9_-]+}/color", handleColorPut).
		Methods(http.MethodPut).
		HeadersRegexp("Content-Type", "application/json")

	r.HandleFunc("/cluster/{name:[a-zA-Z0-9_-]+}", handleClusterDelete).
		Methods(http.MethodDelete)

	r.HandleFunc("/cluster", handleClusterList).
		Methods(http.MethodGet)
}

func handleClusterPost(w http.ResponseWriter, r *http.Request) {

	var postCluster cluster.Cluster
	if err := common.ReadBody(r.Body, &postCluster); err != nil {
		common.WriteError(http.StatusBadRequest, "bad request", w)
		return
	}

	newCluster, err := cluster.GetManagerInstance().RegisterNewCluster(*postCluster.Name, *postCluster.Color)
	if err != nil {
		common.WriteError(http.StatusInternalServerError, "error registering new cluster", w)
	}

	data, err := json.Marshal(newCluster)
	if err != nil {
		common.WriteError(http.StatusInternalServerError, "error preparing response", w)
		return
	}

	common.WriteSuccess(data, w)
}

func handleClusterGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterName := vars["name"]

	cluster, err := cluster.GetManagerInstance().GetCluster(clusterName)
	if err != nil {
		common.WriteError(http.StatusBadRequest, err.Error(), w)
		return
	}

	data, err := json.Marshal(cluster)
	if err != nil {
		common.WriteError(http.StatusInternalServerError, "error preparing response", w)
		return
	}

	common.WriteSuccess(data, w)
}

func handleClusterList(w http.ResponseWriter, r *http.Request) {
	clusters := cluster.GetManagerInstance().GetClusters()
	data, err := json.Marshal(clusters)
	if err != nil {
		common.WriteError(http.StatusInternalServerError, "error preparing response", w)
		return
	}

	common.WriteSuccess(data, w)
}

func handleClusterDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterName := vars["name"]

	_, err := cluster.GetManagerInstance().UnregisterCluster(clusterName)
	if err != nil {
		if err == cluster.ErrFailedDelete {
			common.WriteError(http.StatusInternalServerError, err.Error(), w)
		} else {
			common.WriteError(http.StatusBadRequest, err.Error(), w)
		}
		return
	}
}

// Color Calls

func handleColorPut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterName := vars["name"]

	var colorMessage common.ColorMessage
	if err := common.ReadBody(r.Body, &colorMessage); err != nil {
		common.WriteError(http.StatusBadRequest, "bad request", w)
		return
	}

	cluster, err := cluster.GetManagerInstance().SetClusterColor(clusterName, colorMessage.Color)
	if err != nil {
		common.WriteError(http.StatusNotFound, err.Error(), w)
		return
	}

	data, err := json.Marshal(cluster)
	if err != nil {
		common.WriteError(http.StatusInternalServerError, "error preparing response", w)
		return
	}

	common.WriteSuccess(data, w)
}

func handleColorGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterName := vars["name"]

	cluster, err := cluster.GetManagerInstance().GetCluster(clusterName)
	if err != nil {
		common.WriteError(http.StatusNotFound, err.Error(), w)
		return
	}

	colorMessage := &common.ColorMessage{
		Color: *cluster.Color,
	}

	data, err := json.Marshal(colorMessage)
	if err != nil {
		common.WriteError(http.StatusInternalServerError, "error preparing response", w)
		return
	}

	common.WriteSuccess(data, w)
}
