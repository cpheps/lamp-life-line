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

func TestRegisterNewLamp(t *testing.T) {
	cluster := createTestCluster()

	lamp := cluster.RegisterNewLamp(lampID, listenAddress)

	if *lamp.ID != lampID {
		t.Errorf("Expected ID %s got %s", lampID, *lamp.ID)
	} else if *lamp.ClusterID != clusterID {
		t.Errorf("Expected ClusterID %s got %s", clusterID, *lamp.ClusterID)
	} else if *lamp.ListenAddress != listenAddress {
		t.Errorf("Expected Listen Address %s got %s", listenAddress, *lamp.ListenAddress)
	}

	cacheLamp := cluster.Lamps[lampID]
	if cacheLamp == nil {
		t.Errorf("No lamp found matching id %s", lampID)
	}

	if cacheLamp != lamp {
		t.Errorf("Expecting %v in cache found %v", lamp, cacheLamp)
	}
}

func TestRegisterLamp(t *testing.T) {
	cluster := createTestCluster()
	lamp := createLamp(lampID, clusterID, listenAddress)
	cluster.RegisterLamp(lamp)

	cacheLamp := cluster.Lamps[lampID]
	if cacheLamp == nil {
		t.Errorf("No lamp found matching id %s", lampID)
	}

	if cacheLamp != lamp {
		t.Errorf("Expecting %v in cache found %v", lamp, cacheLamp)
	}
}

func TestUnRegisterLamp(t *testing.T) {
	cluster := createTestCluster()
	lamp := createLamp(lampID, clusterID, listenAddress)

	unregisterLamp, err := cluster.UnRegisterLamp(*lamp.ID)
	if err == nil {
		t.Error("Expecting error from empty cluster when trying to unregister a lamp")
	}

	cluster.RegisterLamp(lamp)

	unregisterLamp, err = cluster.UnRegisterLamp(*lamp.ID)
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	} else if unregisterLamp != lamp {
		t.Errorf("Removed lamp %v expecting %v", unregisterLamp, lamp)
	}

	if _, ok := cluster.Lamps[*lamp.ID]; ok {
		t.Error("Did not remove lamp from cache")
	}

}

func createTestCluster() *Cluster {
	return CreateCluster(&clusterID, &clusterName)
}
