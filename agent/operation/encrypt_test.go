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
		CertFile:   "/Users/kaitoY/Documents/sslcerts/blackbox-agent-cert.pem",
		KeyFile:    "/Users/kaitoY/Documents/sslcerts/private/blackbox-agent-key.pem",
		CaCert:     "/Users/kaitoY/Documents/sslcerts/cacert.pem",
	}
	err := AesEncryptFolderAuto(arg, arg.FolderPath)
	if err != nil {
		fmt.Println(err)
	}
}
