package cluster

import (
	"testing"
)

var testClusterName = "Test"

func TestCreateCluster(t *testing.T) {
	cluster := createTestCluster()

	if *cluster.ID != testClusterID {
		t.Errorf("Expected id %s got %s", testClusterID, *cluster.ID)
	} else if *cluster.Name != testClusterName {
		t.Errorf("Expected name %s got %s", testClusterName, *cluster.Name)
	} else if cluster.Lamps == nil {
		t.Error("Expected Lamps map to be initialized")
	}
}

func TestRegisterNewLamp(t *testing.T) {
	cluster := createTestCluster()
	lamp, err := cluster.RegisterNewLamp(testLampID, testListenAddress)

	if err != nil {
		t.Errorf("Got unexpected error %s", err.Error())
	} else if *lamp.ID != testLampID {
		t.Errorf("Expected ID %s got %s", testLampID, *lamp.ID)
	} else if *lamp.ClusterID != testClusterID {
		t.Errorf("Expected ClusterID %s got %s", testClusterID, *lamp.ClusterID)
	} else if *lamp.ListenAddress != testListenAddress {
		t.Errorf("Expected Listen Address %s got %s", testListenAddress, *lamp.ListenAddress)
	}

	cacheLamp := cluster.Lamps[testLampID]
	if cacheLamp == nil {
		t.Errorf("No lamp found matching id %s", testLampID)
	}

	if cacheLamp != lamp {
		t.Errorf("Expecting %v in cache found %v", lamp, cacheLamp)
	}

	_, err = cluster.RegisterNewLamp(*lamp.ID, *lamp.ListenAddress)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestRegisterLamp(t *testing.T) {
	cluster := createTestCluster()
	lamp := createLamp(testLampID, testClusterID, testListenAddress)
	cluster.RegisterLamp(lamp)

	cacheLamp := cluster.Lamps[testLampID]
	if cacheLamp == nil {
		t.Errorf("No lamp found matching id %s", testLampID)
	}

	if cacheLamp != lamp {
		t.Errorf("Expecting %v in cache found %v", lamp, cacheLamp)
	}
}

func TestUnRegisterLamp(t *testing.T) {
	cluster := createTestCluster()
	lamp := createLamp(testLampID, testClusterID, testListenAddress)

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

func TestGetLamp(t *testing.T) {
	cluster := createTestCluster()
	lamp := createLamp(testLampID, testClusterID, testListenAddress)

	retrieveLamp, err := cluster.GetLamp("nope")
	if err == nil {
		t.Errorf("Expected error")
	}

	cluster.Lamps[*lamp.ID] = lamp
	retrieveLamp, err = cluster.GetLamp(*lamp.ID)
	if err != nil {
		t.Errorf("Got non-nil error: %s", err.Error())
	} else if retrieveLamp != lamp {
		t.Errorf("Expected %v got %v", lamp, retrieveLamp)
	}

}

func createTestCluster() *Cluster {
	return CreateCluster(&testClusterID, &testClusterName)
}
