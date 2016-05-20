package daemon

import (
	"blackbox/cli"
	"fmt"
)

func StartServer(args cli.Args) {
	server := NewServer()
	err := server.Start(args)
	if err != nil {
		fmt.Println("Failed to start server.")
	}
}
