package main

import (
	"fmt"
	"flag"
	"blackbox/cli"
	"blackbox/daemon"
)

func main() {
	args := cli.Parse()
	fmt.Println(args)
	if args.Daemon && args.Port >= 1024 && args.Port <= 65535 {
		daemon.StartServer(args)
	} else {
		flag.PrintDefaults()
	}
}
