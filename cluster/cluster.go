package cluster

//TODO handle updating listenAddress of a registered lamp

import (
	"fmt"
)

//Cluster represents a cluster of Lamps
type Cluster struct {
	ID    *string
	Name  *string
	Lamps map[string]*Lamp
}

//CreateCluster creates a new cluster instance
func CreateCluster(id, name *string) *Cluster {
	return &Cluster{
		ID:    id,
		Name:  name,
		Lamps: make(map[string]*Lamp),
	}
}

//RegisterNewLamp creates a new lamp and registers it to this cluster
func (c *Cluster) RegisterNewLamp(id, listenAddress string) *Lamp {
	fmt.Printf("Registering new Lamp with ID <%s> and ListenAddress <%s> to Cluster <%s>\n", id, listenAddress, *c.Name)

	lamp := createLamp(id, *c.ID, listenAddress)
	c.Lamps[*lamp.ID] = lamp

	fmt.Printf("Registered a new lamp <%s>", *lamp.ID)
	return lamp
}

//RegisterLamp registers an existing lamp to this cluster
func (c *Cluster) RegisterLamp(lamp *Lamp) {
	fmt.Printf("Registering existing Lamp with ID <%s> and ListenAddress <%s> to Cluster <%s>\n", *lamp.ID, *lamp.ListenAddress, *c.Name)
	c.Lamps[*lamp.ID] = lamp
	lamp.ClusterID = c.ID
	fmt.Printf("Registered existing lamp <%s>\n", *lamp.ID)
}

//UnRegisterLamp removes a lamp from the cluster. Returns an error if no lamp was found
func (c *Cluster) UnRegisterLamp(id string) (*Lamp, error) {
	fmt.Printf("Ungregistering Lamp <%s> from Cluster <%s>\n", id, *c.Name)
	lamp := c.Lamps[id]

	if lamp == nil {
		fmt.Printf("Failed to ungregister Lamp <%s> from Cluster <%s>\n", id, *c.Name)
		return nil, fmt.Errorf("No lamp containing id %s in cluster", id)
	}

	lamp.ClusterID = nil

	fmt.Printf("Successfully ungregistered Lamp <%s> from Cluster <%s>\n", id, *c.Name)
	return lamp, nil
}
