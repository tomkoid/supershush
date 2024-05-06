package config

import (
	"errors"
	"fmt"
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
		PollRateMs: 20,
		Mpc:        true,
		PlayerCtl:  true,
	}

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	configFilePath := userConfigDir + "/audiochanger/config.toml"

	configStr, err := readConfigFile(configFilePath)
	if err != nil {
		log.Println("using default config")
		return config
	}

	err = toml.Unmarshal([]byte(configStr), &config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("config: %+v\n", config)

	return config
}
