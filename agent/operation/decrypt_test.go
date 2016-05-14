package operation

import (
	"fmt"
	"testing"
	"blackbox/agent/cli"

)

func TestDecryptFile(t *testing.T) {
	arg := cli.Args {
		Remove 		: true,
		FilePath 	: "/Users/kaitoY/Documents/test/test.txt",
		Suffix 		: "blackbox",
		Target 		: "127.0.0.1:23333",
	}
	err := AesEncryptFileAuto(arg, arg.FilePath)
	arg = cli.Args {
		FilePath 	: "/Users/kaitoY/Documents/test/test.txt.blackbox",
		Suffix 		: "blackbox",
		Target 		: "127.0.0.1:23333",
	}
	err = AesDecryptFileAuto(arg, arg.FilePath)
	fmt.Println(err)
}