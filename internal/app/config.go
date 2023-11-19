package internal

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

const configFileName = ".wipcoconfig"

func SaveAPIKey(apikey string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	configFile := filepath.Join(usr.HomeDir, configFileName)
	err = os.WriteFile(configFile, []byte(apikey), 0600)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadAPIKey() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	configFile := filepath.Join(usr.HomeDir, configFileName)
	data, err := os.ReadFile(configFile)
	if err != nil {
		return ""
	}

	return string(data)
}
