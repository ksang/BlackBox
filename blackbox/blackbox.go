package main

import (
	"blackbox/cli"
	"blackbox/daemon"
	"flag"
)

func main() {
	args := cli.Parse()
	if args.Daemon && args.Port >= 1024 && args.Port <= 65535 {
		daemon.StartServer(args)
	} else {
		flag.PrintDefaults()
	}
}
