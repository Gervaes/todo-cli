package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func getTodos() []Todo {
	projectUrl := getEnvVariable("PROJECT_URL")
	apiKey := getEnvVariable("API_KEY")

	client := &http.Client{}
	url := projectUrl + "/todos?date=eq." + time.Now().Format("2006-01-02")
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", apiKey)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	var todos []Todo
	err = json.Unmarshal(body, &todos)

	if err != nil {
		log.Fatal(err)
	}

	return todos
}
