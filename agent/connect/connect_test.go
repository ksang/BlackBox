package connect

import (
	"fmt"
	"testing"
	"blackbox/worker"
	"blackbox/agent/cli"
)

func TestConnect(t *testing.T) {
	arg := cli.Args{Target: "127.0.0.1:23333"}
	keys := make([]string, 500)
	for i := 0; i < 500; i++ {
		keys[i], _ = worker.GenerateValue(16)
	}
	var res []string
	for _, key := range keys {
		r, _ := RequestSecret(arg, key)
		res = append(res, string(r))
	}
	for _,r := range res {
		fmt.Println(r)
	}
}