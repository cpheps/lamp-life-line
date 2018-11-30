package database

import (
	"errors"
)

// MockDBConnection used for mocking a DBConnection in tests
type MockDBConnection struct {
	GetClusterReturnValue *ClusterModel

	GetAllClustersReturnValue []ClusterModel

	GetClusterReturnErr     bool
	UpdateClusterReturnErr  bool
	GetAllClustersReturnErr bool
	CreateClusterReturnErr  bool
	DeleteClusterReturnErr  bool
}

// GetCluster mocks out geting a cluster. Will return an error if GetClusterReturnErr
// is true.
func (m MockDBConnection) GetCluster(clusterName string) (*ClusterModel, error) {
	if m.GetClusterReturnErr {
		return nil, errors.New("error")
	}

	return m.GetClusterReturnValue, nil
}

// UpdateCluster mocks out dpdating cluster. Will return an error if UpdateClusterReturnErr
// is true.
func (m MockDBConnection) UpdateCluster(cluster *ClusterModel) error {
	if m.UpdateClusterReturnErr {
		return errors.New("error")
	}

	return nil
}

// GetAllClusters mocks out geting all clusters. Will return an error if GetAllClustersReturnErr
// is true.
func (m MockDBConnection) GetAllClusters() ([]ClusterModel, error) {
	if m.GetAllClustersReturnErr {
		return nil, errors.New("error")
	}

	return m.GetAllClustersReturnValue, nil
}

// CreateCluster mocks out creating cluster. Will return an error if CreateClusterReturnErr
// is true.
func (m MockDBConnection) CreateCluster(cluster *ClusterModel) error {
	if m.CreateClusterReturnErr {
		return errors.New("error")
	}

	return nil
}

// DeleteCluster mocks out deleting cluster. Will return an error if DeleteClusterReturnErr
// is true.
func (m MockDBConnection) DeleteCluster(cluster *ClusterModel) error {
	if m.DeleteClusterReturnErr {
		return errors.New("error")
	}

	return nil
}
