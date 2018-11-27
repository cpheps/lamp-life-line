package database

// DBConnection represents a connection to a database
type DBConnection interface {
	GetCluster(clusterName string) (*ClusterModel, error)
	UpdateCluster(cluster *ClusterModel) error
	GetAllClusters() ([]ClusterModel, error)
	CreateCluster(cluster *ClusterModel) error
}
