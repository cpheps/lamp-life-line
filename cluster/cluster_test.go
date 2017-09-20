package cluster

import (
	"testing"
)

var clusterName = "Test"

func TestCreateCluster(t *testing.T) {
	cluster := createTestCluster()

	if *cluster.ID != clusterID {
		t.Errorf("Expected id %s got %s", clusterID, *cluster.ID)
	} else if *cluster.Name != clusterName {
		t.Errorf("Expected name %s got %s", clusterName, *cluster.Name)
	} else if cluster.Lamps == nil {
		t.Error("Expected Lamps map to be initialized")
	}
}

func createTestCluster() *Cluster {
	return CreateCluster(&clusterID, &clusterName)
}
