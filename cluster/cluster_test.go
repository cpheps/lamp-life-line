package cluster

import (
	"reflect"
	"testing"

	"github.com/cpheps/lamp-life-line/database"
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

func Test_Cluster_ToDBModel(t *testing.T) {
	name, color := "cluster", uint32(42)
	cluster := CreateCluster(name, color)

	expected := &database.ClusterModel{
		Name:  name,
		Color: int(color),
	}

	out := cluster.ToDBModel()
	if !reflect.DeepEqual(expected, out) {
		t.Errorf("Expected %+v got %+v", expected, out)
	}
}

func TestFromDBModel(t *testing.T) {
	name, color := "cluster", uint32(42)
	input := &database.ClusterModel{
		Name:  name,
		Color: int(color),
	}

	expected := CreateCluster(name, color)

	out := FromDBModel(input)
	if !reflect.DeepEqual(expected, out) {
		t.Errorf("Expected %+v got %+v", expected, out)
	}
}
