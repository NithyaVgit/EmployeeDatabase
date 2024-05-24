package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var Configs Config

func ReadConfig(configPath string) (bool, error) {
	//Load the configuration file
	yamlFileRead, yamlFileReadErr := ioutil.ReadFile(configPath)
	if yamlFileReadErr != nil {
		return false, yamlFileReadErr
	}

	// Assign the configutation details into the package variable
	unmarshalYamlFileErr := yaml.Unmarshal(yamlFileRead, &Configs)
	if unmarshalYamlFileErr != nil {
		return false, unmarshalYamlFileErr
	}

	return true, nil
}
