package daemon

import (
	"fmt"
	"blackbox/cli"
)

func StartServer(args cli.Args) {
	server := NewServer()
	err := server.Start(args)
	if err != nil {
		fmt.Println("Failed to start server.")
	}
}