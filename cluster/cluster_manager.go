package cluster

import (
	"fmt"
	"sync"

	"github.com/satori/go.uuid"
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
	clusterID := generateUUID(m.clusterCache)
	cluster := CreateCluster(&clusterID, &name, &color)
	m.clusterCache[clusterID] = cluster

	fmt.Printf("Registered Cluster <%s> with assigned id <%s>\n", name, clusterID)
	return cluster
}

//GetCluster retrieves a cluster with a given id
func (m *Manager) GetCluster(clusterID string) (*Cluster, error) {
	cluster := m.clusterCache[clusterID]

	if cluster == nil {
		return nil, fmt.Errorf("No cluster with id <%s> found", clusterID)
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
func (m *Manager) UnregisterCluster(id string) (*Cluster, error) {
	fmt.Printf("Unregistering new Cluster with id <%s>\n", id)
	cluster := m.clusterCache[id]

	if cluster == nil {
		fmt.Printf("Failed to ungregister Cluster with id <%sn", id)
		return nil, fmt.Errorf("No Cluster with id <%s> in cache", id)
	}

	delete(m.clusterCache, id)

	fmt.Printf("Unregistered Cluster <%s> with assigned id <%s>\n", *cluster.Name, id)

	return cluster, nil
}

func generateUUID(clusterCache map[string]*Cluster) string {
	id := uuid.NewV4().String()
	for ; clusterCache[id] != nil; id = uuid.NewV4().String() {
	}
	return id
}
