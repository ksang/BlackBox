package operation

import (
	"blackbox/agent/cli"
	"log"
	"os"
	"path/filepath"
)

type PlainFile struct {
	Path string
	MD5  string
}

type PlainFolder struct {
	Files []PlainFile
}

func fileOp(f func(cli.Args, string) error,
	arg cli.Args,
	path string,
	errc chan error) {
	err := f(arg, path)
	errc <- err
}

// method for encrypt/decrypt folder concurrently
func folderOp(f func(cli.Args, string) error,
	arg cli.Args,
	path string) error {
	errc := make(chan error)
	var done int
	files, err := getFolderFiles(path)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, path := range files {
		path := path
		go fileOp(f, arg, path, errc)
	}
	for {
		select {
		case err := <-errc:
			if err != nil {
				log.Println(err)
			}
			done++
		default:
			// all task finished
			if done == len(files) {
				return nil
			}
		}
	}
}

func getFolderFiles(root string) ([]string, error) {
	_, err := os.Stat(root)
	if err != nil {
		return nil, err
	}
	var files []string
	err = filepath.Walk(root,
		func(path string, f os.FileInfo, err error) error {
			if err != nil {
				log.Print(err)
				return nil
			}
			if !(f.IsDir()) {
				//log.Print(path)
				files = append(files, path)
			}
			return nil
		})
	return files, nil
}
