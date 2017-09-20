package cluster

import (
	"testing"
)

var (
	lampID        = "MAC"
	clusterID     = "UUID"
	listenAddress = "192.168.1.1"
)

func TestCreateLamp(t *testing.T) {
	lamp := createLamp(lampID, clusterID, listenAddress)

	if *lamp.ID != lampID {
		t.Errorf("Expected ID %s got %s", lampID, *lamp.ID)
	} else if *lamp.ClusterID != clusterID {
		t.Errorf("Expected ClusterID %s got %s", clusterID, *lamp.ClusterID)
	} else if *lamp.ListenAddress != listenAddress {
		t.Errorf("Expected Listen Address %s got %s", listenAddress, *lamp.ListenAddress)
	}
}
