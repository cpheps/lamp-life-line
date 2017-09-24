package cluster

//TODO: Add color change functionality and logging

//Lamp is a single lamp within a cluster
type Lamp struct {
	ID            *string
	ClusterID     *string
	ListenAddress *string
	//need some variable for color, will hold until color api is explored better
}

//createLamp creates a new lamp object
func createLamp(id, clusterID, listenAddress string) *Lamp {
	return &Lamp{
		ID:            &id,
		ClusterID:     &clusterID,
		ListenAddress: &listenAddress,
	}
}
