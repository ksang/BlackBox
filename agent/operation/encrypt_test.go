package operation

import (
	"blackbox/agent/cli"
	"fmt"
	"testing"
)

func TestEncryptFolder(t *testing.T) {
	arg := cli.Args{
		Remove:     true,
		FolderPath: "/Users/kaitoY/Documents/test/tests",
		Suffix:     "blackbox",
		Target:     "127.0.0.1:23333",
	}
	err := AesEncryptFolderAuto(arg, arg.FolderPath)
	if err != nil {
		fmt.Println(err)
	}
}
