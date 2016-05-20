package operation

import (
	"blackbox/agent/cli"
	"fmt"
	"testing"
)

func TestDecryptFile(t *testing.T) {
	arg := cli.Args{
		Remove:   true,
		FilePath: "/Users/kaitoY/Documents/test/test.txt",
		Suffix:   "blackbox",
		Target:   "127.0.0.1:23333",
	}
	err := AesEncryptFileAuto(arg, arg.FilePath)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	arg = cli.Args{
		FilePath: "/Users/kaitoY/Documents/test/test.txt.blackbox",
		Suffix:   "blackbox",
		Target:   "127.0.0.1:23333",
	}
	err = AesDecryptFileAuto(arg, arg.FilePath)
	if err != nil {
		fmt.Println("Error", err)
	}
}

func TestDecryptFolder(t *testing.T) {
	arg := cli.Args{
		FolderPath: "/Users/kaitoY/Documents/test/tests_de",
		Suffix:     "blackbox",
		Target:     "127.0.0.1:23333",
	}
	err := AesDecryptFolderAuto(arg, arg.FolderPath)
	if err != nil {
		fmt.Println(err)
	}
}
