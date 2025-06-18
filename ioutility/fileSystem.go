package ioutility

import (
	"os"
	"path/filepath"

	"github.com/pezhmankasraee/pklog/v2"
)

const permission os.FileMode = 0755

/*
File permission is 0755
*/
func MakeDirectory(path string) {
	err := os.MkdirAll(path, permission)
	if err != nil {
		pklog.CreateLog(pklog.FatalError, err.Error())
	}
	pklog.CreateLog(pklog.Information, "path to "+path+" was created successfully")
}

func CheckPathExist(path string) {

	_, err := os.Stat(path)
	if err != nil {
		pklog.CreateLog(pklog.Warning, "base path does not exist")
		MakeDirectory(path)
	} else {
		pklog.CreateLog(pklog.Information, "base path exists")
	}
}

// readFile opens the file at the given path and returns it.
// The caller is responsible for closing the returned file.
func readFile(pathToFile string) (file *os.File, e error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		pklog.CreateLog(pklog.Error, "file cannot be open: "+err.Error())
		return nil, err
	}
	return file, nil
}

func splitPath(fullPath string) (directory string, filename string) {

	dir := filepath.Dir(fullPath)
	file := filepath.Base(fullPath)

	return dir, file
}
