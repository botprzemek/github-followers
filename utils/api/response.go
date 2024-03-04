package api

import (
	"io"
	"log"
	"net/http"
	"os"
)

func parseBody(body io.ReadCloser) ([]byte, error) {
	buffer, err := io.ReadAll(body)

	if err != nil {
		os.Exit(0)
	}

	return buffer, err
}

func Request(method string, url string, headers *http.Header) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal("[SECRETS] failed to load configuration file")
	}

	req.Header = *headers

	res, err := client.Do(req)
	if err != nil {
		os.Exit(0)
	}

	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			log.Fatal("[SECRETS] failed to load configuration file")
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		log.Fatal("[SECRETS] failed to load configuration file")
	}

	buffer, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("[SECRETS] failed to load configuration file")
	}

	println(string(buffer))

	return string(buffer), err
}
