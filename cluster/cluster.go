package cluster

import (
	"github.com/cpheps/lamp-life-line/database"
)

//Cluster represents a cluster of Lamps
type Cluster struct {
	Name  *string `json:"name"`
	Color *uint32 `json:"color"`
}

//CreateCluster creates a new cluster instance
func CreateCluster(name string, color uint32) *Cluster {
	return &Cluster{
		Name:  &name,
		Color: &color,
	}
}

// FromDBModel create a cluster from a Database Model
func FromDBModel(dbCluster *database.ClusterModel) *Cluster {
	clusterColor := uint32(dbCluster.Color)
	return &Cluster{
		Name:  &dbCluster.Name,
		Color: &clusterColor,
	}
}

// ToDBModel convert a cluster to a Database Model
func (c Cluster) ToDBModel() *database.ClusterModel {
	return &database.ClusterModel{
		Name:  *c.Name,
		Color: int(*c.Color),
	}
}
