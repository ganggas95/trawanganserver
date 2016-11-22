package job

import (
	"fmt"
	"os"
	"strings"
)

func RenameFile(file *os.File, namaFile string) (string, error) {
	file.Close()
	oldPath := file.Name()
	path := strings.Split(oldPath, "/")
	tempPath := strings.Split(path[1], ".")
	fmt.Println(file.Name())
	fmt.Println(path[0] + "/" + namaFile + "." + tempPath[1])
	newPath := path[0] + "/" + namaFile + "." + tempPath[1]
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return "", err
	}
	return namaFile + "." + tempPath[1], nil

}
