package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

type User struct {
	Login string `json:"login"`
}

func CreateFile(fileName string) error {
	log.Println("[CACHE] creating new cache file")

	file, err := os.Create(fileName)
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal("[CACHE] failed to create cache file, make sure you have write permission to this directory")
		}
	}(file)

	return err
}

func LoadFile(fileName string) ([]byte, error) {
	log.Println("[CACHE] loading cache file")

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("[CACHE] cache file not found")

		err = CreateFile(fileName)
		return nil, err
	}

	if file == nil {
		log.Fatal("[CACHE] failed to load cache file")
	}

	return file, nil
}

func LoadConfiguration(fileName string) {
	err := godotenv.Load(fileName)
	if err != nil {
		_ = CreateFile(fileName)
		log.Fatal("[SECRETS] failed to load configuration file")
	}

	token = os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("[SECRETS] token not found, add TOKEN key to configuration file")
	}
}

func SaveCache(fileName string) error {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("[CACHE] cache file not found")
	}

	if bytes == nil {
		err = CreateFile(fileName)
		return err
	}

	return err
}

func Default(*cli.Context) error {
	file, err := LoadFile(cacheFileName)
	if err != nil {
		return err
	}

	var data []User
	err = json.Unmarshal(file, &data)
	if err != nil {
		return err
	}

	println(data[0].Login)

	err = SaveCache(cacheFileName)
	if err != nil {
		return err
	}

	return err
}

const configurationFileName = ".env"
const cacheFileName = "cache.json"

var token string

func main() {
	LoadConfiguration(configurationFileName)

	app := &cli.App{
		Name:                 "github-followers",
		Usage:                "make an explosive entrance",
		EnableBashCompletion: true,
		Action:               Default,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// https://api.github.com/user/followers
// https://api.github.com/user/following/cherrymui
