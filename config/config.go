package config

import (
	"github-followers/utils/file"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

const fileName = ".env"

var config = map[string]string{
	"token": "",
	"test":  "",
}

func Load() {
	if err := godotenv.Load(fileName); err != nil {
		log.Println("[SECRETS] failed to load configuration file")

		if err = file.Create(fileName); err != nil {
			log.Fatal("[SECRETS] failed to load configuration file")
		}
	}

	for key := range config {
		config[key] = os.Getenv(strings.ToUpper(key))

		if config[key] != "" {
			continue
		}

		log.Println("[SECRETS] " + key + " not found, add " + strings.ToUpper(key) + " value to configuration file")
	}
}

func Get(key string) string {
	return config[key]
}
