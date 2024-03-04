package api

import (
	"github-followers/config"
	"log"
	"net/http"
)

func GetUsers() {
	token := "Bearer " + config.Get("token")
	headers := http.Header{}

	headers.Set("Accept", "application/vnd.github+json")
	headers.Set("Authorization", token)
	headers.Set("X-GitHub-Api-Version", "2022-11-28")

	_, err := Request("GET", "https://api.github.com/user/followers", &headers)
	if err != nil {
		log.Fatal("[SECRETS] failed to load configuration file")
	}
}
