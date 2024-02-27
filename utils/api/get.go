package api

import (
	"fmt"
	"github-followers/config"
	"io"
	"log"
	"net/http"
)

func GetUsers() {
	token := "Bearer " + config.Get("token")

	request, err := http.NewRequest("GET", "https://api.github.com/user/followers", nil)
	if err != nil {
		return
	}

	request.Header.Add("Accept", "application/vnd.github+json")
	request.Header.Add("Authorization", token)
	request.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Fatal("T")
		}
	}(response.Body)

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		fmt.Println("Request failed with status code:", response.StatusCode)
		return
	}

	fmt.Println("Response:", response.Body)
}

//func Request() {
//	req, err := http.NewRequest("GET", "https://api.knuryknurow.pl/v1/players", nil)
//
//	req.Header.Add("x-api-key", "5994859fb1d4ca2ddae414ee488d7f57")
//
//	if err != nil {
//		os.Exit(0)
//	}
//
//	client := &http.Client{}
//
//	res, err := client.Do(req)
//
//	if err != nil {
//		os.Exit(1)
//	}
//
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//			os.Exit(0)
//		}
//	}(res.Body)
//
//	if res.StatusCode != http.StatusOK {
//		fmt.Println(res.StatusCode)
//		os.Exit(0)
//	}
//
//	buffer, err := io.ReadAll(res.Body)
//
//	if err != nil {
//		fmt.Println("Test 5")
//		os.Exit(0)
//	}
//
//	fmt.Printf("Data: %v\n", string(buffer))
//}
