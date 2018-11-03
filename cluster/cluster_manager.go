package cluster

import (
	"fmt"
	"errors"
	"sync"
)

//Manager handles cluster lifecycle
type Manager struct {
	clusterCache map[string]*Cluster
	mutex        *sync.RWMutex
}

var instance *Manager
var once sync.Once

//GetManagerInstance returns singleton ClusterManager
func GetManagerInstance() *Manager {
	once.Do(func() {
		instance = createManager()
	})

	return instance
}

func createManager() *Manager {
	return &Manager{
		clusterCache: make(map[string]*Cluster),
		mutex:        &sync.RWMutex{},
	}
}

//RegisterNewCluster creates a new cluster and registers to manager
func (m *Manager) RegisterNewCluster(name string, color uint32) *Cluster {
	fmt.Printf("Registering new Cluster with name <%s>\n", name)
	cluster := CreateCluster(name, color)
	m.clusterCache[name] = cluster

	fmt.Println("Registered Cluster <",name, ">")
	return cluster
}

//GetCluster retrieves a cluster with a given id
func (m *Manager) GetCluster(clusterName string) (*Cluster, error) {
	cluster := m.clusterCache[clusterName]

	if cluster == nil {
		return nil, fmt.Errorf("No cluster with id <%s> found", clusterName)
	}

	return cluster, nil
}

//GetClusters returns all clusters being managed
func (m *Manager) GetClusters() []*Cluster {
	clusters := make([]*Cluster, 0, len(m.clusterCache))
	for _, cluster := range m.clusterCache {
		clusters = append(clusters, cluster)
	}

	return clusters
}

//UnregisterCluster removes the cluster form the manager.
//Returns an error if no cluster is found
func (m *Manager) UnregisterCluster(name string) (*Cluster, error) {
	fmt.Println("Unregistering new Cluster with id <", name, ">")
	cluster := m.clusterCache[name]

	if cluster == nil {
		fmt.Println("Failed to unregister Cluster with id <", name, ">")
		return nil, errors.New("no such cluster")
	}

	delete(m.clusterCache, name)

	fmt.Println("Unregistered Cluster <", name, ">")

	return cluster, nil

}