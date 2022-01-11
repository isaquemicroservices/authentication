package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Init variable to struct configuration
var config *Configuration

// Load tries to read from config.json file a valid JSON with all settings
func Load() {
	var pathConfig string = "/etc/ms-auth/config.json"
	dataPath, err := ioutil.ReadFile(pathConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(dataPath, &config); err != nil {
		log.Fatal(err)
	}
}

// Get get data of configurations
func Get() *Configuration {
	return config
}
