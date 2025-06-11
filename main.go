package main

import (
	"flag"

	"github.com/pezhmankasraee/pklog/v2"
	"github.com/pezhmankasraee/pksetdev/config"
	"github.com/pezhmankasraee/pksetdev/help"
	"github.com/pezhmankasraee/pksetdev/ioutility"
)

var (
	pathToConfigFile string
	isHelp           bool
	isVersion        bool
)

func main() {

	flag.Parse()
	help.ShowHelp(isHelp)
	help.ShowVersion(isVersion)

	pklog.CreateLog(pklog.Information, "pksetdev starts ...")
	pklog.CreateLog(pklog.Information, "Yaml file locations: "+pathToConfigFile)

	yamlFile := ioutility.ReadYamlFile(pathToConfigFile)

	pklog.CreateLog(pklog.Information, "basepath: "+yamlFile.BasePath)
}

func init() {
	flag.StringVar(&pathToConfigFile, "p", config.PathToDefaultConfigYamlFile, "Path to config file")
	flag.StringVar(&pathToConfigFile, "path", config.PathToDefaultConfigYamlFile, "Path to config file")
	flag.BoolVar(&isHelp, "h", false, "show this help")
	flag.BoolVar(&isHelp, "help", false, "show this help")
	flag.BoolVar(&isVersion, "v", false, "show the version of this application")
	flag.BoolVar(&isVersion, "version", false, "show the version of this application")
}
