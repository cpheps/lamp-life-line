package main

import "fmt"

//Version current version of the Lamp Life Line server
var Version string

//BuildTime build time of binary
var BuildTime string

func main() {
	fmt.Printf("Running Lamp Life Line version %s build on %s", Version, BuildTime)
}

//Minimum Handlers to Launch
//Register Cluster
//Register Lamp
//Change Color
