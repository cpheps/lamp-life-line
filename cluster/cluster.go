package cluster

//TODO handle updating listenAddress of a registered lamp

import (
	"fmt"
)

//Cluster represents a cluster of Lamps
type Cluster struct {
	ID    *string          `json:"id"`
	Name  *string          `json:"name"`
	Lamps map[string]*Lamp `json:"lamps"`
	//TODO Current cluster color
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
func (c *Cluster) RegisterNewLamp(id, listenAddress string) (*Lamp, error) {
	fmt.Printf("Registering new Lamp with ID <%s> and ListenAddress <%s> to Cluster <%s>\n", id, listenAddress, *c.Name)
	if _, ok := c.Lamps[id]; ok {
		return nil, fmt.Errorf("Lamp with ID <%s> already exists in cluster <%s>", id, *c.ID)
	}

	lamp := createLamp(id, *c.ID, listenAddress)
	c.Lamps[*lamp.ID] = lamp
	//TODO set color of lamp

	fmt.Printf("Registered a new lamp <%s>", *lamp.ID)
	return lamp, nil
}

//RegisterLamp registers an existing lamp to this cluster
func (c *Cluster) RegisterLamp(lamp *Lamp) error {
	fmt.Printf("Registering existing Lamp with ID <%s> and ListenAddress <%s> to Cluster <%s>\n", *lamp.ID, *lamp.ListenAddress, *c.Name)
	if _, ok := c.Lamps[*lamp.ID]; ok {
		return fmt.Errorf("A lamp with ID <%s> already exists in cluster <%s>", *lamp.ID, *c.ID)
	}

	c.Lamps[*lamp.ID] = lamp
	lamp.ClusterID = c.ID
	fmt.Printf("Registered existing lamp <%s>\n", *lamp.ID)
	return nil
}

//UnRegisterLamp removes a lamp from the cluster. Returns an error if no lamp was found
func (c *Cluster) UnRegisterLamp(id string) (*Lamp, error) {
	fmt.Printf("Unregistering Lamp <%s> from Cluster <%s>\n", id, *c.Name)
	lamp := c.Lamps[id]

	if lamp == nil {
		fmt.Printf("Failed to unregister Lamp <%s> from Cluster <%s>\n", id, *c.Name)
		return nil, fmt.Errorf("No lamp with id <%s> in cluster", id)
	}

	delete(c.Lamps, id)
	lamp.ClusterID = nil

	fmt.Printf("Successfully unregistered Lamp <%s> from Cluster <%s>\n", id, *c.Name)
	return lamp, nil
}

//GetLamp retrieves a lamp with a given id
func (c *Cluster) GetLamp(id string) (*Lamp, error) {
	lamp := c.Lamps[id]

	if lamp == nil {
		return nil, fmt.Errorf("No Lamp with id <%s> found in cluster <%s>", id, *c.Name)
	}

	return lamp, nil
}
