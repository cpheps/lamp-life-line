package main

import (
	"log"
	"os"

	"github.com/cpheps/lamp-life-line/cluster"
	"github.com/cpheps/lamp-life-line/database"
	"github.com/cpheps/lamp-life-line/server"
)

func main() {
	log.Printf("Running Lamp Life Line")

	// Load database connection
	dbConn, err := database.NewConnection()
	if err != nil {
		log.Printf("Unable to connect to Database: %s", err.Error())
		os.Exit(1)
	}

	// Set the connection in the manager
	cluster.GetManagerInstance().SetDBConnection(dbConn)

	errors := server.StartServer()

	err = <-errors
	if err != nil {
		log.Printf("Error running Lamp Life Line: %s", err.Error())
	}
}
