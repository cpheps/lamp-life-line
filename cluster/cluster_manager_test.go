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
	testClusterName, color := "name", int32(4)
	manager := GetManagerInstance()
	clearManager()

	cluster := manager.RegisterNewCluster(testClusterName, color)

	if *cluster.Name != testClusterName {
		t.Errorf("Expected cluster with name %s got %s", testClusterName, *cluster.Name)
	} else if *cluster.Color != color {
		t.Errorf("Expected cluster with color %d got %d", color, *cluster.Color)
	}

	if cacheCluster, ok := manager.clusterCache[*cluster.ID]; ok {
		if cacheCluster != cluster {
			t.Error("Manager stored wrong cluster in cache for ID")
		}
	} else {
		t.Error("No cluster found in cache for ID")
	}
}

func TestGetCluster(t *testing.T) {
	manager := GetManagerInstance()
	clearManager()
	id, name, color := "id", "cluster", int32(42)
	cluster := CreateCluster(&id, &name, &color)

	_, err := manager.GetCluster(*cluster.ID)

	if err == nil {
		t.Error("Expected err for missing Cluster in Cache")
	}

	manager.clusterCache[*cluster.ID] = cluster

	getCluster, err := manager.GetCluster(*cluster.ID)

	if err != nil {
		t.Error("Unexpected error thrown for getting known cluster")
	} else if getCluster != cluster {
		t.Errorf("Got cluster %v expected %v", getCluster, cluster)
	}
}

func TestUnregisterCluster(t *testing.T) {
	manager := GetManagerInstance()
	clearManager()

	id, name, color := "id", "cluster", int32(42)
	cluster := CreateCluster(&id, &name, &color)

	_, err := manager.UnregisterCluster(*cluster.ID)

	if err == nil {
		t.Error("Expected err for missing Cluster in Cache")
	}

	manager.clusterCache[*cluster.ID] = cluster

	getCluster, err := manager.UnregisterCluster(*cluster.ID)

	if err != nil {
		t.Error("Unexpected error thrown for getting known cluster")
	} else if getCluster != cluster {
		t.Errorf("Got cluster %v expected %v", getCluster, cluster)
	}

	if _, ok := manager.clusterCache[*cluster.ID]; ok {
		t.Error("Cluster was not removed from cache")
	}

}

func TestGenerateUUID(t *testing.T) {
	testCache := make(map[string]*Cluster)

	for i := 0; i < 100; i++ {
		uuid := generateUUID(testCache)
		if _, ok := testCache[uuid]; ok {
			t.Errorf("Duplicate %s UUID generated", uuid)
			break
		}
		testCache[uuid] = nil
	}
}

func TestGetClusters(t *testing.T) {
	manager := GetManagerInstance()
	clearManager()

	idOne, idTwo, name, color := "id", "idTwo", "cluster", int32(42)
	clusterOne := CreateCluster(&idOne, &name, &color)
	clusterTwo := CreateCluster(&idTwo, &name, &color)

	manager.clusterCache[*clusterOne.ID] = clusterOne
	manager.clusterCache[*clusterTwo.ID] = clusterTwo

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
