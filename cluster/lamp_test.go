package cluster

import (
	"testing"
)

func TestCreateLamp(t *testing.T) {
	id, clusterID, listenAddress := "MAC", "UUID", "192.168.1.1"

	lamp := createLamp(id, clusterID, listenAddress)

	if *lamp.ID != id {
		t.Errorf("Expected ID %s got %s", id, *lamp.ID)
	} else if *lamp.ClusterID != clusterID {
		t.Errorf("Expected ClusterID %s got %s", clusterID, *lamp.ClusterID)
	} else if *lamp.ListenAddress != listenAddress {
		t.Errorf("Expected Listen Address %s got %s", listenAddress, *lamp.ListenAddress)
	}
}