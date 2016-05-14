package operation

import (
	"log"
	"os"
	"path/filepath"
)

type PlainFile struct {
	Path string
	MD5 string
}

type PlainFolder struct {
	Files []PlainFile
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

