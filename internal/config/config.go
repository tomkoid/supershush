package config

import (
	"errors"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func configFileExists(configFilePath string) bool {
	// if path exists
	_, err := os.Stat(configFilePath)
	if err != nil {
		return false
	}

	return true
}

func readConfigFile(configFilePath string) (string, error) {
	if !configFileExists(configFilePath) {
		return "", errors.New("config file does not exist")
	} else {
		file, err := os.ReadFile(configFilePath)
		if err != nil {
			return "", err
		}

		configStr := string(file)

		return configStr, nil
	}
}

func GetConfig() Config {
	// default config
	var config Config = Config{
		Mpc:       false,
		PlayerCtl: true,
		Resume:    false,
	}

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	configFilePath := userConfigDir + "/supershush/config.toml"

	configStr, err := readConfigFile(configFilePath)
	if err != nil {
		log.Println("Using default config.")
		if err.Error() != "config file does not exist" {
			log.Printf("config error: %s\n", err.Error())
		}
		return config
	} else {
		log.Printf("Using config from %s.\n", configFilePath)
	}

	err = toml.Unmarshal([]byte(configStr), &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
