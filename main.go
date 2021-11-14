package main

import (
	"log"

	"paper.com/paper/config"
)

func checkFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Setup Config
	conf, err := config.New()
	checkFatalError(err)
	// Print Config
	conf.Print()
	// Init Server
	server := SetupServer(conf)
	// Start Server
	server.Start()
}
