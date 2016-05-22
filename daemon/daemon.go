package daemon

import (
	"blackbox/cli"
	"log"
)

func StartServer(args cli.Args) {
	server := NewServer()
	err := server.Start(args)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
