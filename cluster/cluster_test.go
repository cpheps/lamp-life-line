package cluster

import (
	"testing"
)

var testClusterName = "Test"

func TestCreateCluster(t *testing.T) {
	name, color := "cluster", uint32(42)
	cluster := CreateCluster(name, color)

	if *cluster.Name != name {
		t.Errorf("Expected name %s got %s", name, *cluster.Name)
	} else if *cluster.Color != color {
		t.Errorf("Expected color %d got %d", color, *cluster.Color)
	}
}
