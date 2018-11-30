package cluster

import (
	"errors"
	"log"
	"sync"

	"github.com/cpheps/lamp-life-line/database"
)

//Manager handles cluster lifecycle
type Manager struct {
	clusterCache map[string]*Cluster
	mutex        *sync.RWMutex
	conn         database.DBConnection
}

// ErrFailedDelete represents a Database failure to delete a cluster
var ErrFailedDelete = errors.New("unable to delete cluster")

var instance *Manager
var once sync.Once

//GetManagerInstance returns singleton Manager
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

// SetDBConnection sets the database connection for this manager
func (m *Manager) SetDBConnection(conn database.DBConnection) {
	m.conn = conn

	// load all models as clusers from the database
	models, err := m.conn.GetAllClusters()
	if err != nil {
		log.Printf("Failed to load all Models from Database: %s", err.Error())
		return
	}

	for _, model := range models {
		cluster := FromDBModel(&model)
		m.clusterCache[*cluster.Name] = cluster
	}
}

//RegisterNewCluster creates a new cluster and registers to manager
func (m *Manager) RegisterNewCluster(name string, color uint32) (*Cluster, error) {
	log.Printf("Registering new Cluster with name <%s>", name)
	cluster := CreateCluster(name, color)

	if err := m.conn.CreateCluster(cluster.ToDBModel()); err != nil {
		log.Printf("Failed to save to Database: %s", err.Error())
		return nil, err
	}

	// Don't cache unless we can put in database
	m.clusterCache[name] = cluster

	log.Printf("Registered Cluster <%s>", name)
	return cluster, nil
}

//GetCluster retrieves a cluster with a given id
func (m *Manager) GetCluster(clusterName string) (*Cluster, error) {
	cluster, ok := m.clusterCache[clusterName]
	if !ok {
		model, err := m.conn.GetCluster(clusterName)
		if err != nil {
			log.Printf("Failed to lookup cluster in Database: %s", err.Error())
			return nil, errors.New("unable to find cluster")
		}

		cluster = FromDBModel(model)
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
		log.Printf("Failed to unregister Cluster <%s>", name)
		return nil, errors.New("no such cluster")
	}

	if err := m.conn.DeleteCluster(cluster.ToDBModel()); err != nil {
		log.Printf("Failed to remove cluster from Database: %s", err.Error())
		return nil, errors.New("unable to delete cluster")
	}

	// Do not remove from cache until we delete from database
	delete(m.clusterCache, name)

	log.Printf("Unregistered Cluster <%s>", name)

	return cluster, nil

}

// SetClusterColor sets the color of a cluster
func (m *Manager) SetClusterColor(name string, color uint32) (*Cluster, error) {
	cluster, err := m.GetCluster(name)
	if err != nil {
		return nil, err
	}

	cluster.Color = &color

	err = m.conn.UpdateCluster(cluster.ToDBModel())
	// The cluster will still have the color until sotred in memory so don't return an error
	// but rather log
	if err != nil {
		log.Printf("Failed to update Cluster <%s> color in Database", *cluster.Name)
	}

	return cluster, nil
}
