package file

import (
	"log"
	"os"
)

func Create(fileName string) error {
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

func Load(fileName string) ([]byte, error) {
	log.Println("[CACHE] loading cache file")

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("[CACHE] cache file not found")

		err = Create(fileName)
		return nil, err
	}

	if file == nil {
		log.Fatal("[CACHE] failed to load cache file")
	}

	return file, nil
}
