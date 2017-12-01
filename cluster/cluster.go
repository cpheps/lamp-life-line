package cluster

//Cluster represents a cluster of Lamps
type Cluster struct {
	ID    *string `json:"id"`
	Name  *string `json:"name"`
	Color *int32  `json:"color"`
}

//CreateCluster creates a new cluster instance
func CreateCluster(id, name *string, color *int32) *Cluster {
	return &Cluster{
		ID:    id,
		Name:  name,
		Color: color,
	}
}
