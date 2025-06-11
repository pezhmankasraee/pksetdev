package main

import (
	"flag"

	"github.com/pezhmankasraee/pklog/v2"
	"github.com/pezhmankasraee/pksetdev/config"
	"github.com/pezhmankasraee/pksetdev/help"
	"github.com/pezhmankasraee/pksetdev/ioutility"
	pkgmanagement "github.com/pezhmankasraee/pksetdev/pkgManagement"
)

<<<<<<< HEAD
<<<<<<< HEAD
=======

>>>>>>> 22081c8 (add show version functionality for help)
var (
	pathToConfigFile string
	isHelp           bool
	isVersion        bool
<<<<<<< HEAD
)
=======
var pathToConfigFile string
var isHelp bool
var isInstall bool
>>>>>>> c05d5e6 (add downloader)
=======
	isInstall bool
)
>>>>>>> 22081c8 (add show version functionality for help)

func main() {

	flag.Parse()
	help.ShowHelp(isHelp)
	help.ShowVersion(isVersion)

	pklog.CreateLog(pklog.Information, "pksetdev starts ...")
	pklog.CreateLog(pklog.Information, "Yaml file locations: "+pathToConfigFile)

	yamlFile := ioutility.ReadYamlFile(pathToConfigFile)

	pklog.CreateLog(pklog.Information, "base path in yaml: "+yamlFile.BasePath)

	ioutility.CheckPathExist(yamlFile.BasePath)

	pkgmanagement.InstallApplications(yamlFile, isInstall)
}

func init() {
	flag.BoolVar(&isInstall, "i", false, "Download and install the applications")
	flag.BoolVar(&isInstall, "install", false, "Download and install the applications")
	flag.StringVar(&pathToConfigFile, "p", config.PathToDefaultConfigYamlFile, "Path to config file")
	flag.StringVar(&pathToConfigFile, "path", config.PathToDefaultConfigYamlFile, "Path to config file")
<<<<<<< HEAD
<<<<<<< HEAD
	flag.BoolVar(&isHelp, "h", false, "show this help")
	flag.BoolVar(&isHelp, "help", false, "show this help")
	flag.BoolVar(&isVersion, "v", false, "show the version of this application")
	flag.BoolVar(&isVersion, "version", false, "show the version of this application")
=======
	flag.BoolVar(&isHelp, "h", false, "Show this help")
	flag.BoolVar(&isHelp, "help", false, "Show this help")
>>>>>>> c05d5e6 (add downloader)
=======
	flag.BoolVar(&isHelp, "h", false, "Show this help")
	flag.BoolVar(&isHelp, "help", false, "Show this help")
=======
	flag.BoolVar(&isHelp, "h", false, "show this help")
	flag.BoolVar(&isHelp, "help", false, "show this help")
	flag.BoolVar(&isVersion, "v", false, "show the version of this application")
	flag.BoolVar(&isVersion, "version", false, "show the version of this application")
>>>>>>> f2ba19d (add show version functionality for help)
>>>>>>> 22081c8 (add show version functionality for help)
}
