package config

import (
	"github.com/pokekrishna/kome-automation/filesystem"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func GetConf(filePath string) (map[interface{}]interface{}, error) {
	// validate config file path
	if err := filesystem.FileExists(filePath); err != nil {
		log.Fatal("Error validating Config File: ", err)
	}

	return loadYAML(filePath)
}

func loadYAML(filePath string) (map[interface{}]interface{}, error) {
	// Read content from file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		// if there is a problem return with the error
		return nil, err
	}

	// load content to yaml
	m := make(map[interface{}]interface{})
	if err := yaml.Unmarshal(content, &m); err != nil {
		return nil, err
	}
	return m, nil
}
