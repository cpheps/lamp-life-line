package cluster

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	manager := GetManagerInstance()

	secondManager := GetManagerInstance()

	if manager != secondManager {
		t.Error("Returned two separate pointesr for singleton")
	}
}

func TestRegisterNewCluster(t *testing.T) {
	testClusterName, color := "name", uint32(4)
	manager := GetManagerInstance()
	clearManager()

	cluster := manager.RegisterNewCluster(testClusterName, color)

	if *cluster.Name != testClusterName {
		t.Errorf("Expected cluster with name %s got %s", testClusterName, *cluster.Name)
	} else if *cluster.Color != color {
		t.Errorf("Expected cluster with color %d got %d", color, *cluster.Color)
	}

	if cacheCluster, ok := manager.clusterCache[*cluster.Name]; ok {
		if cacheCluster != cluster {
			t.Error("Manager stored wrong cluster in cache for Name")
		}
	} else {
		t.Error("No cluster found in cache for ID")
	}
}

func TestGetCluster(t *testing.T) {
	manager := GetManagerInstance()
	clearManager()
	name, color := "cluster", uint32(42)
	cluster := CreateCluster(name, color)

	_, err := manager.GetCluster(*cluster.Name)

	if err == nil {
		t.Error("Expected err for missing Cluster in Cache")
	}

	manager.clusterCache[*cluster.Name] = cluster

	getCluster, err := manager.GetCluster(*cluster.Name)

	if err != nil {
		t.Error("Unexpected error thrown for getting known cluster")
	} else if getCluster != cluster {
		t.Errorf("Got cluster %v expected %v", getCluster, cluster)
	}
}

func TestUnregisterCluster(t *testing.T) {
	manager := GetManagerInstance()
	clearManager()

	name, color := "cluster", uint32(42)
	cluster := CreateCluster(name, color)

	_, err := manager.UnregisterCluster(*cluster.Name)

	if err == nil {
		t.Error("Expected err for missing Cluster in Cache")
	}

	manager.clusterCache[*cluster.Name] = cluster

	getCluster, err := manager.UnregisterCluster(*cluster.Name)

	if err != nil {
		t.Error("Unexpected error thrown for getting known cluster")
	} else if getCluster != cluster {
		t.Errorf("Got cluster %v expected %v", getCluster, cluster)
	}

	if _, ok := manager.clusterCache[*cluster.Name]; ok {
		t.Error("Cluster was not removed from cache")
	}

}


func TestGetClusters(t *testing.T) {
	manager := GetManagerInstance()
	clearManager()

	nameOne, nameTwo, color := "One", "Two", uint32(42)
	clusterOne := CreateCluster(nameOne, color)
	clusterTwo := CreateCluster(nameTwo, color)

	manager.clusterCache[*clusterOne.Name] = clusterOne
	manager.clusterCache[*clusterTwo.Name] = clusterTwo

	clusters := manager.GetClusters()

	if length := len(clusters); length != 2 {
		t.Errorf("Expected 2 clusters got %d", length)
	}
}

func clearManager() {
	manager := GetManagerInstance()

	for k := range manager.clusterCache {
		delete(manager.clusterCache, k)
	}
}
