package pkgmanagement

import (
	"fmt"
	"os"
	"strings"

	"github.com/pezhmankasraee/pklog/v2"
	"github.com/pezhmankasraee/pksetdev/ioutility"
	"github.com/pezhmankasraee/pksetdev/model"
	"github.com/pezhmankasraee/pksetdev/net"
)

func InstallApplications(yaml *model.YamlFile, isInstall bool) {

	if isInstall {

		for _, value := range yaml.Applications {
			install(&value, yaml.BasePath)
		}
	}
}

func install(application *model.Application, basePath string) {

	appLowerName := strings.ToLower(application.Name)
	fmt.Println("")
	pklog.CreateLog(pklog.Information, fmt.Sprintf("start installing [   %s   ] ...", appLowerName))

	applicationBasePath := basePath + "/" + application.Name
	_, err := os.Stat(applicationBasePath)
	if err == nil {
		os.RemoveAll(applicationBasePath)
	}

	ioutility.MakeDirectory(applicationBasePath)

	net.Download(application, applicationBasePath)

	if err := uncompressFiles(application, applicationBasePath); err != nil {
		pklog.CreateLog(pklog.FatalError, err.Error())
	}
}

func uncompressFiles(application *model.Application, applicationBasePath string) error {
	destination := applicationBasePath + "/" + application.Filename

	if isTarBallFile(application.Filename) {
		pklog.CreateLog(pklog.Information, fmt.Sprintf("extracting .tar.gz: %s", destination))
		if err := ioutility.ExtractTarGz(destination); err != nil {
			return fmt.Errorf(".tar.gz file: %s", err.Error())
		}
		pklog.CreateLog(pklog.Information, "extracted successfully.")
	} else if strings.HasSuffix(application.Filename, ".zip") {
		pklog.CreateLog(pklog.Information, fmt.Sprintf("extracting .zip: %s", destination))
		if err := ioutility.ExtractZip(destination); err != nil {
			return fmt.Errorf(".zip file: %s", err.Error())
		}
		pklog.CreateLog(pklog.Information, "extracted successfully.")
	} else {
		return fmt.Errorf("unrecognizable compressed file")
	}

	return nil
}

func isTarBallFile(fileName string) bool {
	return strings.HasSuffix(fileName, ".tar.gz") ||
		strings.HasSuffix(fileName, ".tgz")
}
