package ioutility

import (
	"os"

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
