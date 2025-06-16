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
	pklog.CreateLog(pklog.Information, "start installing [   "+appLowerName+"   ] ...")

	applicationBasePath := basePath + "/" + application.Name
	_, err := os.Stat(applicationBasePath)
	if err != nil {
		pklog.CreateLog(pklog.Warning, "path to "+applicationBasePath+" does not exist")
		ioutility.MakeDirectory(applicationBasePath)
	}
	net.Download(application, applicationBasePath)

}
