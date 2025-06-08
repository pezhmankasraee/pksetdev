package ioutility

import (
	"os"

	"github.com/pezhmankasraee/pklog/v2"
	"gopkg.in/yaml.v3"

	"github.com/pezhmankasraee/pksetdev/model"
)

func ReadYamlFile(yamlFileLocation string) *model.YamlFile {
	configfile, err := os.ReadFile(yamlFileLocation)
	if err != nil {
		pklog.CreateLog(pklog.FatalError, "yaml file"+yamlFileLocation+" is not accessible")
	}

	return unmarshalYamlFile(configfile)
}

func unmarshalYamlFile(configFile []byte) *model.YamlFile {

	var setup model.YamlFile
	if err := yaml.Unmarshal(configFile, &setup); err != nil {
		pklog.CreateLog(pklog.Error, "config yaml file cannot be unmashalised")
	}

	return &setup
}
