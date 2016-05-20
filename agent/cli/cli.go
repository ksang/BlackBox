package cli

import (
	"errors"
	"flag"
)

const (
	MODE_ENCRYPT = 0
	MODE_DECRYPT = 1
)

type Args struct {
	Mode       int
	Remove     bool
	FilePath   string
	FolderPath string
	Suffix     string
	Target     string
	CertFile   string
	KeyFile    string
}

func Parse() (Args, error) {
	var encrypt = flag.Bool("e", false, "Encrypt mode.")
	var decrypt = flag.Bool("d", false, "Decrypt mode.")
	var remove = flag.Bool("r", false, "Remove the original file.(encrypt only)")
	var filepath = flag.String("f", "", "Path to the file.")
	var folderpath = flag.String("p", "", "Path to the folder.")
	var suffix = flag.String("s", "blackbox", "Encrypt file suffix.")
	var target = flag.String("t",
		"127.0.0.1:23333",
		"Target address with port.")
	var certFile = flag.String("c", "agent.pem", "Agent cert file.")
	var keyFile = flag.String("k", "agent.key", "Agent key file.")
	flag.Parse()
	mode := -1
	if *encrypt && !*decrypt {
		mode = MODE_ENCRYPT
	} else if !*encrypt && *decrypt {
		mode = MODE_DECRYPT
	} else {
		return Args{}, errors.New("You must select one of encrypt or decrypt mode.")
	}
	if len(*filepath) == 0 && len(*folderpath) == 0 {
		return Args{}, errors.New("You must provide filepath or folderpath.")
	}
	return Args{
		Mode:       mode,
		FilePath:   *filepath,
		FolderPath: *folderpath,
		Suffix:     *suffix,
		Remove:     *remove,
		Target:     *target,
		CertFile:   *certFile,
		KeyFile:    *keyFile,
	}, nil
}
