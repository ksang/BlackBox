/*
Blackbox server is the manager for encryption keys.
It stores keys in a key-value persistent cache system.
Agent will use file hash value as key and server respond with a randomly generated value.
The value is used for encrypt/decrypt files.
*/
package main

import (
	"flag"
	"github.com/ksang/blackbox/cli"
	"github.com/ksang/blackbox/daemon"
)

func main() {
	args := cli.Parse()
	if args.Daemon && args.Port >= 1024 && args.Port <= 65535 {
		daemon.StartServer(args)
	} else {
		flag.PrintDefaults()
	}
}
