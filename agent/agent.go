package main

import (
    "fmt"
    "flag"
    _ "blackbox/agent/connect"
    "blackbox/agent/cli"
)


func main() {
    args, err := cli.Parse()
    if err != nil {
    	fmt.Println(err)
    	flag.PrintDefaults()
    	return
    }
    fmt.Println(args)
}