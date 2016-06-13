/*
Blackbox agent is a client side program which encrypt/decrypt files.
It request blackbox server for encryption keys when need to do operations.
Agent never keeps the key.
*/
package main

import (
	"flag"
	"github.com/ksang/blackbox/agent/cli"
	"github.com/ksang/blackbox/agent/operation"
	"log"
)

func main() {
	args, err := cli.Parse()
	if err != nil {
		log.Print(err)
		flag.PrintDefaults()
		return
	}
	switch args.Mode {
	case cli.MODE_ENCRYPT:
		if len(args.FilePath) > 0 {
			err = operation.AesEncryptFileAuto(args, args.FilePath)
		} else if len(args.FolderPath) > 0 {
			err = operation.AesEncryptFolderAuto(args, args.FolderPath)
		}
		if err != nil {
			log.Fatal("FAILED:", err)
			return
		}
	case cli.MODE_DECRYPT:
		if len(args.FilePath) > 0 {
			err = operation.AesDecryptFileAuto(args, args.FilePath)
		} else if len(args.FolderPath) > 0 {
			err = operation.AesDecryptFolderAuto(args, args.FolderPath)
		}
		if err != nil {
			log.Fatal("FAILED:", err)
			return
		}
	}
}
