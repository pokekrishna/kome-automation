package config

import (
	"github.com/pokekrishna/kome-automation/filesystem"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)


type Config struct {
	ServerConfig ServerConfig `yaml:"server_config"`
}

type ServerConfig struct {
	BindIP string `yaml:"bind_ip"`
	BindPort int `yaml:"bind_port"`
}

func LoadConf(filePath string)  (*Config, error) {
	// validate config file path
	if err := filesystem.FileExists(filePath); err != nil {
		log.Fatal("Error validating Config File: ", err)
	}
	return loadFile(filePath)
}

func loadFile(filePath string) (*Config, error) {
	// Read content from file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		// if there is a problem return with the error
		return nil, err
	}

	// load content to yaml
	m := &Config{}
	if err := yaml.Unmarshal(content, m); err != nil {
		return nil, err
	}
	return m, nil
}
