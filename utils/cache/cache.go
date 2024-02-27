package cache

import (
	"encoding/json"
	"github-followers/models"
	"github-followers/utils/file"
	"log"
	"os"
)

const fileName = "cache.json"

func Save() error {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("[CACHE] cache file not found")
	}

	if bytes == nil {
		err = file.Create(fileName)
		return err
	}

	return err
}

func Load() error {
	cache, err := file.Load(fileName)
	if err != nil {
		return err
	}

	var data models.Cache
	if err = json.Unmarshal(cache, &data); err != nil {
		return err
	}

	println(data.GetUsers())

	return err
}
