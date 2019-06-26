package kits

import (
	"os"
	"os/exec"
	// "log"
	"path/filepath"
	"strings"
)

func GetAppPath() string{
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	// log.Println("GetAppPath========>>")
	index := strings.LastIndex(path,string(os.PathSeparator))
	// log.Println(path[:index])
	return path[:index]
}