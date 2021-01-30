package config

import (
	"github.com/pokekrishna/kome-automation/filesystem"
	"log"
)

func ValidateConfFile(filePath string) error {
	// validate config file path
	if err := filesystem.FileExists(filePath); err != nil{
		log.Fatal("Error validating Config File: ", err)
	}

	// is file parsable to yaml ?
	// TODO: bookmark
}

