package cluster

//Cluster represents a cluster of Lamps
type Cluster struct {
	ID    *string `json:"id"`
	Name  *string `json:"name"`
	Color *uint32 `json:"color"`
}

//CreateCluster creates a new cluster instance
func CreateCluster(id, name *string, color *uint32) *Cluster {
	return &Cluster{
		ID:    id,
		Name:  name,
		Color: color,
	}
}
