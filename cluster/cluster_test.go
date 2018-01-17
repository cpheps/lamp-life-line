package cluster

import (
	"testing"
)

var testClusterName = "Test"

func TestCreateCluster(t *testing.T) {
	id, name, color := "id", "cluster", uint32(42)
	cluster := CreateCluster(&id, &name, &color)

	if *cluster.ID != id {
		t.Errorf("Expected id %s got %s", id, *cluster.ID)
	} else if *cluster.Name != name {
		t.Errorf("Expected name %s got %s", name, *cluster.Name)
	} else if *cluster.Color != color {
		t.Errorf("Expected color %d got %d", color, *cluster.Color)
	}
}
