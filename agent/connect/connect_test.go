package connect

import (
	"blackbox/agent/cli"
	"blackbox/worker"
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	arg := cli.Args{Target: "127.0.0.1:23333"}
	keys := make([]string, 10)
	for i := 0; i < 10; i++ {
		key, err := worker.GenerateValue(16)
		if err != nil {
			fmt.Println("Error GenerateValue(): ", err)
			continue
		}
		keys[i] = key
	}
	fmt.Println("Keys:", len(keys))
	var res []string
	for _, key := range keys {
		r, err := RequestSecret(arg, key)
		if err != nil {
			fmt.Println("RequestSecret Error:", err)
			continue
		}
		//fmt.Println("Got secret pair:", r)
		res = append(res, string(r))
	}
	for _, r := range res {
		fmt.Println(r)
	}
	//re-check old keys
	for _, key := range keys {
		r, err := RequestSecret(arg, key)
		if err != nil {
			fmt.Println("RequestSecret Error:", err)
			continue
		}
		fmt.Println("Old Key Received:", r)
	}
}
