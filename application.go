package main

import (
	"fmt"

	"github.com/cpheps/lamp-life-line/server"
)

//Version current version of the Lamp Life Line server
var Version string

//BuildTime build time of binary
var BuildTime string

func main() {
	fmt.Printf("Running Lamp Life Line version %s build on %s\n", Version, BuildTime)

	errors := server.StartServer()

	err := <-errors
	if err != nil {
		fmt.Printf("Error running Lamp Life Line: %s\n", err.Error())
	}
}

//Minimum Handlers to Launch
//Register Cluster
//Register Lamp
//Change Color
