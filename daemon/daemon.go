/*
Package daemon implements blackbox server listening service.
It will listen on the address/port and accept for incoming connections.
It also creates workers goroutines and waiting for imcoming requests.
The workers will act similar to a thread pool, ensure stable performance.
*/
package daemon

import (
	"blackbox/cli"
	"log"
)

// Start server
func StartServer(args cli.Args) {
	server := NewServer()
	err := server.Start(args)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
