package cluster

import (
	"errors"
	"log"
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
	log.Printf("Registering new Cluster with name <%s>", name)
	cluster := CreateCluster(name, color)
	m.clusterCache[name] = cluster

	log.Printf("Registered Cluster <%s>", name)
	return cluster
}

//GetCluster retrieves a cluster with a given id
func (m *Manager) GetCluster(clusterName string) (*Cluster, error) {
	cluster := m.clusterCache[clusterName]

	if cluster == nil {
		return nil, errors.New("no cluster found")
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
	log.Printf("Unregistering new Cluster with name <%s>", name)
	cluster := m.clusterCache[name]

	if cluster == nil {
		log.Printf("Failed to unregister Cluster with id <%s>", name)
		return nil, errors.New("no such cluster")
	}

	delete(m.clusterCache, name)

	log.Printf("Unregistered Cluster <%s>", name)

	return cluster, nil

}
