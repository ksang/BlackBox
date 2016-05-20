package main

import (
	"blackbox/cli"
	"blackbox/daemon"
	"flag"
	"fmt"
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
