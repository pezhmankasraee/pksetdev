package model

type YamlFile struct {
	BasePath     string        `yaml:"basePath"`
	Applications []Application `yaml:"applications"`
}

type Application struct {
	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	Url       string `yaml:"url"`
	Filename  string `yaml:"filename"`
	Hash      string `yaml:"hash"`
	Algorithm string `yaml:"algorithm"`
	OverWrite bool   `yaml:"overWrite"`
}
