package operation

import (
	"fmt"
	"testing"

)

func TestGetFolderFile(t *testing.T) {
	files, err := getFolderFiles("/Users/kaitoY/Documents/")
	fmt.Println(err)
	fmt.Println(len(files))
	for _,v := range files {
		fmt.Println(v)
	}
}