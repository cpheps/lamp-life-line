package cluster

import (
	"testing"
)

var (
	testLampID        = "MAC"
	testClusterID     = "UUID"
	testListenAddress = "192.168.1.1"
)

func TestCreateLamp(t *testing.T) {
	lamp := createLamp(testLampID, testClusterID, testListenAddress)

	if *lamp.ID != testLampID {
		t.Errorf("Expected ID %s got %s", testLampID, *lamp.ID)
	} else if *lamp.ClusterID != testClusterID {
		t.Errorf("Expected ClusterID %s got %s", testClusterID, *lamp.ClusterID)
	} else if *lamp.ListenAddress != testListenAddress {
		t.Errorf("Expected Listen Address %s got %s", testListenAddress, *lamp.ListenAddress)
	}
}
