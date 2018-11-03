package cluster

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
