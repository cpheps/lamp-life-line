package cluster

import (
	"reflect"
	"testing"

	"github.com/cpheps/lamp-life-line/database"
)

func TestGetInstance(t *testing.T) {
	manager := GetManagerInstance()

	secondManager := GetManagerInstance()

	if manager != secondManager {
		t.Error("Returned two separate pointesr for singleton")
	}
}

func TestSetDBConnection_Error(t *testing.T) {
	manager := GetManagerInstance()
	clearManager()

	manager.SetDBConnection(&database.MockDBConnection{
		GetAllClustersReturnValue: []database.ClusterModel{
			{
				Name:  "one",
				Color: 1,
			},
		},
		GetAllClustersReturnErr: true,
	})

	if len(manager.clusterCache) != 0 {
		t.Error("manager is not empty on db fail")
	}
}

func TestSetDBConnection_NoError(t *testing.T) {
	manager := GetManagerInstance()
	clearManager()

	manager.SetDBConnection(&database.MockDBConnection{
		GetAllClustersReturnValue: []database.ClusterModel{
			{
				Name:  "one",
				Color: 1,
			},
		},
		GetAllClustersReturnErr: false,
	})

	if len(manager.clusterCache) != 1 {
		t.Errorf("Cache size of %d expected 1", len(manager.clusterCache))
	}
}

func TestRegisterNewCluster_NoError(t *testing.T) {
	testClusterName, color := "name", uint32(4)
	manager := GetManagerInstance()
	clearManager()

	manager.SetDBConnection(&database.MockDBConnection{})

	cluster, err := manager.RegisterNewCluster(testClusterName, color)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

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

func TestRegisterNewCluster_Error(t *testing.T) {
	testClusterName, color := "name", uint32(4)
	manager := GetManagerInstance()
	clearManager()

	manager.SetDBConnection(&database.MockDBConnection{
		CreateClusterReturnErr: true,
	})

	_, err := manager.RegisterNewCluster(testClusterName, color)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestGetCluster_NoError_InCache(t *testing.T) {
	manager := GetManagerInstance()
	manager.SetDBConnection(&database.MockDBConnection{})
	clearManager()

	name, color := "cluster", uint32(42)
	cluster := CreateCluster(name, color)

	manager.clusterCache[*cluster.Name] = cluster

	getCluster, err := manager.GetCluster(*cluster.Name)

	if err != nil {
		t.Error("Unexpected error thrown for getting known cluster")
	} else if getCluster != cluster {
		t.Errorf("Got cluster %v expected %v", getCluster, cluster)
	}
}

func TestGetCluster_Error(t *testing.T) {
	manager := GetManagerInstance()
	manager.SetDBConnection(&database.MockDBConnection{
		GetClusterReturnErr: true,
	})
	clearManager()

	_, err := manager.GetCluster("test")
	if err == nil {
		t.Error("Expected err for missing Cluster in Cache")
	}
}

func TestGetCluster_NoError_InDB(t *testing.T) {
	manager := GetManagerInstance()
	manager.SetDBConnection(&database.MockDBConnection{
		GetClusterReturnValue: &database.ClusterModel{
			Name:  "cluster",
			Color: 42,
		},
	})
	clearManager()

	name, color := "cluster", uint32(42)
	expected := CreateCluster(name, color)

	getCluster, err := manager.GetCluster(name)

	if err != nil {
		t.Error("Unexpected error thrown for getting known cluster")
	} else if !reflect.DeepEqual(expected, getCluster) {
		t.Errorf("Got cluster %v expected %v", getCluster, expected)
	}
}

func TestUnregisterCluster_NoDBError(t *testing.T) {
	manager := GetManagerInstance()
	manager.SetDBConnection(&database.MockDBConnection{})
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

func TestUnregisterCluster_DBError(t *testing.T) {
	manager := GetManagerInstance()
	manager.SetDBConnection(&database.MockDBConnection{
		DeleteClusterReturnErr: true,
	})
	clearManager()

	name, color := "cluster", uint32(42)
	cluster := CreateCluster(name, color)
	manager.clusterCache[*cluster.Name] = cluster

	_, err := manager.UnregisterCluster(name)

	if err == nil {
		t.Error("Expected err for missing Cluster in Cache")
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

func TestSetClusterColor_NoDBError(t *testing.T) {
	manager := GetManagerInstance()
	manager.SetDBConnection(&database.MockDBConnection{})
	clearManager()

	name, color := "cluster", uint32(42)
	cluster := CreateCluster(name, color)
	manager.clusterCache[*cluster.Name] = cluster

	expectedColor := uint32(4)
	out, err := manager.SetClusterColor(name, expectedColor)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if *out.Color != expectedColor {
		t.Errorf("Expected color %d got %d", expectedColor, *out.Color)
	}
}

func TestSetClusterColor_MissingCluster(t *testing.T) {
	manager := GetManagerInstance()
	manager.SetDBConnection(&database.MockDBConnection{
		GetClusterReturnErr: true,
	})
	clearManager()

	_, err := manager.SetClusterColor("test", uint32(1))
	if err == nil {
		t.Error("Expected error")
	}
}
