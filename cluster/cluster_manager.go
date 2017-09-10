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

//GetInstance returns singleton ClusterManager
func GetInstance() *Manager {
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
func (m *Manager) RegisterNewCluster(name string) *Cluster {
	fmt.Printf("Registering new Cluster with name <%s>", name)
	clusterID := generateUUID(m.clusterCache)
	cluster := CreateCluster(&clusterID, &name)
	m.clusterCache[clusterID] = cluster

	fmt.Printf("Registered Cluster <%s> with assigned id <%s>", name, clusterID)
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

//TODO add remove function

func generateUUID(clusterCache map[string]*Cluster) string {
	id := uuid.NewV4().String()
	for ; clusterCache[id] != nil; id = uuid.NewV4().String() {
	}
	return id
}