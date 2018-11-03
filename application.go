package main

import (
	"log"

	"github.com/cpheps/lamp-life-line/server"
)

func main() {
	log.Printf("Running Lamp Life Line")

	errors := server.StartServer()

	err := <-errors
	if err != nil {
		log.Printf("Error running Lamp Life Line: %s", err.Error())
	}
}
