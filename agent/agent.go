package main

import (
	"blackbox/agent/cli"
	"blackbox/agent/operation"
	"flag"
	"fmt"
)

func main() {
	args, err := cli.Parse()
	if err != nil {
		fmt.Println(err)
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
			fmt.Println("FAILED:", err)
			return
		}
	case cli.MODE_DECRYPT:
		if len(args.FilePath) > 0 {
			err = operation.AesDecryptFileAuto(args, args.FilePath)
		} else if len(args.FolderPath) > 0 {
			err = operation.AesDecryptFolderAuto(args, args.FolderPath)
		}
		if err != nil {
			fmt.Println("FAILED:", err)
			return
		}
	}
}
