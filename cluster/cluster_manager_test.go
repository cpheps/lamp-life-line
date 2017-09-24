package cluster

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	manager := GetInstance()

	secondManager := GetInstance()

	if manager != secondManager {
		t.Error("Returned two separate pointesr for singleton")
	}
}

func TestRegisterNewCluster(t *testing.T) {
	manager := GetInstance()
	clearManager()

	cluster := manager.RegisterNewCluster(testClusterName)

	if *cluster.Name != testClusterName {
		t.Errorf("Expected cluster with name %s got %s", testClusterName, *cluster.Name)
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
	manager := GetInstance()
	clearManager()
	cluster := createTestCluster()

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
	manager := GetInstance()
	clearManager()
	cluster := createTestCluster()

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

func clearManager() {
	manager := GetInstance()

	for k := range manager.clusterCache {
		delete(manager.clusterCache, k)
	}
}
